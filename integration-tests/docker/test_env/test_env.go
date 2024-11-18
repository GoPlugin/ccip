package test_env

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	tc "github.com/testcontainers/testcontainers-go"

	"github.com/goplugin/plugin-testing-framework/lib/blockchain"
	ctf_config "github.com/goplugin/plugin-testing-framework/lib/config"
	"github.com/goplugin/plugin-testing-framework/lib/docker"
	"github.com/goplugin/plugin-testing-framework/lib/docker/test_env"
	"github.com/goplugin/plugin-testing-framework/lib/logging"
	"github.com/goplugin/plugin-testing-framework/lib/logstream"
	"github.com/goplugin/plugin-testing-framework/lib/utils/runid"
	"github.com/goplugin/pluginv3.0/v2/core/services/plugin"

	d "github.com/goplugin/pluginv3.0/integration-tests/docker"
)

var (
	ErrFundCLNode = "failed to fund CL node"
)

type CLClusterTestEnv struct {
	Cfg           *TestEnvConfig
	DockerNetwork *tc.DockerNetwork
	LogStream     *logstream.LogStream
	TestConfig    ctf_config.GlobalTestConfig

	/* components */
	ClCluster              *ClCluster
	MockAdapter            *test_env.Killgrave
	PrivateEthereumConfigs []*ctf_config.EthereumNetworkConfig
	EVMNetworks            []*blockchain.EVMNetwork
	rpcProviders           map[int64]*test_env.RpcProvider
	l                      zerolog.Logger
	t                      *testing.T
	isSimulatedNetwork     bool
}

func NewTestEnv() (*CLClusterTestEnv, error) {
	log.Logger = logging.GetLogger(nil, "CORE_DOCKER_ENV_LOG_LEVEL")
	network, err := docker.CreateNetwork(log.Logger)
	if err != nil {
		return nil, err
	}
	return &CLClusterTestEnv{
		DockerNetwork: network,
		l:             log.Logger,
	}, nil
}

// WithTestEnvConfig sets the test environment cfg.
// Sets up private ethereum chain and MockAdapter containers with the provided cfg.
func (te *CLClusterTestEnv) WithTestEnvConfig(cfg *TestEnvConfig) *CLClusterTestEnv {
	te.Cfg = cfg
	if cfg.MockAdapter.ContainerName != "" {
		n := []string{te.DockerNetwork.Name}
		te.MockAdapter = test_env.NewKillgrave(n, te.Cfg.MockAdapter.ImpostersPath, test_env.WithContainerName(te.Cfg.MockAdapter.ContainerName), test_env.WithLogStream(te.LogStream))
	}
	return te
}

func (te *CLClusterTestEnv) WithTestInstance(t *testing.T) *CLClusterTestEnv {
	te.t = t
	te.l = logging.GetTestLogger(t)
	if te.MockAdapter != nil {
		te.MockAdapter.WithTestInstance(t)
	}
	return te
}

func (te *CLClusterTestEnv) StartEthereumNetwork(cfg *ctf_config.EthereumNetworkConfig) (blockchain.EVMNetwork, test_env.RpcProvider, error) {
	// if environment is being restored from a previous state, use the existing config
	// this might fail terribly if temporary folders with chain data on the host machine were removed
	if te.Cfg != nil && te.Cfg.EthereumNetworkConfig != nil {
		cfg = te.Cfg.EthereumNetworkConfig
	}

	te.l.Info().
		Str("Execution Layer", string(*cfg.ExecutionLayer)).
		Str("Ethereum Version", string(*cfg.EthereumVersion)).
		Str("Custom Docker Images", fmt.Sprintf("%v", cfg.CustomDockerImages)).
		Msg("Starting Ethereum network")

	builder := test_env.NewEthereumNetworkBuilder()
	c, err := builder.WithExistingConfig(*cfg).
		WithTest(te.t).
		WithLogStream(te.LogStream).
		Build()
	if err != nil {
		return blockchain.EVMNetwork{}, test_env.RpcProvider{}, err
	}

	n, rpc, err := c.Start()

	if err != nil {
		return blockchain.EVMNetwork{}, test_env.RpcProvider{}, err
	}

	return n, rpc, nil
}

func (te *CLClusterTestEnv) StartMockAdapter() error {
	return te.MockAdapter.StartContainer()
}

func (te *CLClusterTestEnv) StartClCluster(nodeConfig *plugin.Config, count int, secretsConfig string, testconfig ctf_config.GlobalTestConfig, opts ...ClNodeOption) error {
	if te.Cfg != nil && te.Cfg.ClCluster != nil {
		te.ClCluster = te.Cfg.ClCluster
	} else {
		// prepend the postgres version option from the toml config
		if testconfig.GetPluginImageConfig().PostgresVersion != nil && *testconfig.GetPluginImageConfig().PostgresVersion != "" {
			opts = append([]func(c *ClNode){
				func(c *ClNode) {
					c.PostgresDb.EnvComponent.ContainerVersion = *testconfig.GetPluginImageConfig().PostgresVersion
				},
			}, opts...)
		}
		opts = append(opts, WithSecrets(secretsConfig))
		te.ClCluster = &ClCluster{}
		for i := 0; i < count; i++ {
			ocrNode, err := NewClNode([]string{te.DockerNetwork.Name}, *testconfig.GetPluginImageConfig().Image, *testconfig.GetPluginImageConfig().Version, nodeConfig, te.LogStream, opts...)
			if err != nil {
				return err
			}
			te.ClCluster.Nodes = append(te.ClCluster.Nodes, ocrNode)
		}
	}

	// Set test logger
	if te.t != nil {
		for _, n := range te.ClCluster.Nodes {
			n.SetTestLogger(te.t)
		}
	}

	// Start/attach node containers
	return te.ClCluster.Start()
}

func (te *CLClusterTestEnv) Terminate() error {
	// TESTCONTAINERS_RYUK_DISABLED=false by default so ryuk will remove all
	// the containers and the Network
	return nil
}

type CleanupOpts struct {
	TestName string
}

// Cleanup cleans the environment up after it's done being used, mainly for returning funds when on live networks and logs.
func (te *CLClusterTestEnv) Cleanup(opts CleanupOpts) error {
	te.l.Info().Msg("Cleaning up test environment")

	runIdErr := runid.RemoveLocalRunId(te.TestConfig.GetLoggingConfig().RunId)
	if runIdErr != nil {
		te.l.Warn().Msgf("Failed to remove .run.id file due to: %s (not a big deal, you can still remove it manually)", runIdErr.Error())
	}

	if te.t == nil {
		return fmt.Errorf("cannot cleanup test environment without a testing.T")
	}

	if te.ClCluster == nil || len(te.ClCluster.Nodes) == 0 {
		return fmt.Errorf("plugin nodes are nil, unable to cleanup plugin nodes")
	}

	te.logWhetherAllContainersAreRunning()

	err := te.handleNodeCoverageReports(opts.TestName)
	if err != nil {
		te.l.Error().Err(err).Msg("Error handling node coverage reports")
	}

	return nil
}

// handleNodeCoverageReports handles the coverage reports for the plugin nodes
func (te *CLClusterTestEnv) handleNodeCoverageReports(testName string) error {
	testName = strings.ReplaceAll(testName, "/", "_")
	showHTMLCoverageReport := te.TestConfig.GetLoggingConfig().ShowHTMLCoverageReport != nil && *te.TestConfig.GetLoggingConfig().ShowHTMLCoverageReport
	isCI := os.Getenv("CI") != ""

	te.l.Info().
		Bool("showCoverageReportFlag", showHTMLCoverageReport).
		Bool("isCI", isCI).
		Bool("show", showHTMLCoverageReport || isCI).
		Msg("Checking if coverage report should be shown")

	var covHelper *d.NodeCoverageHelper

	if showHTMLCoverageReport || isCI {
		// Stop all nodes in the plugin cluster.
		// This is needed to get go coverage profile from the node containers https://go.dev/doc/build-cover#FAQ
		// TODO: fix this as it results in: ERR LOG AFTER TEST ENDED ... INF 🐳 Stopping container
		err := te.ClCluster.Stop()
		if err != nil {
			return err
		}

		clDir, err := getPluginDir()
		if err != nil {
			return err
		}

		var coverageRootDir string
		if os.Getenv("GO_COVERAGE_DEST_DIR") != "" {
			coverageRootDir = filepath.Join(os.Getenv("GO_COVERAGE_DEST_DIR"), testName)
		} else {
			coverageRootDir = filepath.Join(clDir, ".covdata", testName)
		}

		var containers []tc.Container
		for _, node := range te.ClCluster.Nodes {
			containers = append(containers, node.Container)
		}

		covHelper, err = d.NewNodeCoverageHelper(context.Background(), containers, clDir, coverageRootDir)
		if err != nil {
			return err
		}
	}

	// Show html coverage report when flag is set (local runs)
	if showHTMLCoverageReport {
		path, err := covHelper.SaveMergedHTMLReport()
		if err != nil {
			return err
		}
		te.l.Info().Str("testName", testName).Str("filePath", path).Msg("Plugin node coverage html report saved")
	}

	// Save percentage coverage report when running in CI
	if isCI {
		// Save coverage percentage to a file to show in the CI
		path, err := covHelper.SaveMergedCoveragePercentage()
		if err != nil {
			te.l.Error().Err(err).Str("testName", testName).Msg("Failed to save coverage percentage for test")
		} else {
			te.l.Info().Str("testName", testName).Str("filePath", path).Msg("Plugin node coverage percentage report saved")
		}
	}

	return nil
}

// getPluginDir returns the path to the plugin directory
func getPluginDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("cannot determine the path of the calling file")
	}
	dir := filepath.Dir(filename)
	pluginDir := filepath.Clean(filepath.Join(dir, "../../.."))
	return pluginDir, nil
}

func (te *CLClusterTestEnv) logWhetherAllContainersAreRunning() {
	for _, node := range te.ClCluster.Nodes {
		if node.Container == nil {
			continue
		}

		isCLRunning := node.Container.IsRunning()
		isDBRunning := node.PostgresDb.Container.IsRunning()

		if !isCLRunning {
			te.l.Warn().Str("Node", node.ContainerName).Msg("Plugin node was not running, when test ended")
		}

		if !isDBRunning {
			te.l.Warn().Str("Node", node.ContainerName).Msg("Postgres DB is not running, when test ended")
		}
	}
}

func (te *CLClusterTestEnv) GetRpcProvider(chainId int64) (*test_env.RpcProvider, error) {
	if rpc, ok := te.rpcProviders[chainId]; ok {
		return rpc, nil
	}

	return nil, fmt.Errorf("no RPC provider available for chain ID %d", chainId)
}

func (te *CLClusterTestEnv) GetFirstEvmNetwork() (*blockchain.EVMNetwork, error) {
	if len(te.EVMNetworks) == 0 {
		return nil, fmt.Errorf("no EVM networks available")
	}

	return te.EVMNetworks[0], nil
}

func (te *CLClusterTestEnv) GetEVMNetworkForChainId(chainId int64) (*blockchain.EVMNetwork, error) {
	for _, network := range te.EVMNetworks {
		if network.ChainID == chainId {
			return network, nil
		}
	}

	return nil, fmt.Errorf("no EVM network available for chain ID %d", chainId)
}