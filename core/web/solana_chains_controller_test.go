package web_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/manyminds/api2go/jsonapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commoncfg "github.com/goplugin/plugin-common/pkg/config"
	"github.com/goplugin/plugin-common/pkg/types"
	solcfg "github.com/goplugin/plugin-solana/pkg/solana/config"

	"github.com/goplugin/pluginv3.0/v2/core/internal/cltest"
	"github.com/goplugin/pluginv3.0/v2/core/internal/testutils"
	"github.com/goplugin/pluginv3.0/v2/core/internal/testutils/configtest"
	"github.com/goplugin/pluginv3.0/v2/core/services/plugin"
	"github.com/goplugin/pluginv3.0/v2/core/web"
	"github.com/goplugin/pluginv3.0/v2/core/web/presenters"
)

func Test_SolanaChainsController_Show(t *testing.T) {
	t.Parallel()

	const validId = "Plugin-12"

	testCases := []struct {
		name           string
		inputId        string
		wantStatusCode int
		want           func(t *testing.T, app *cltest.TestApplication) *types.ChainStatus
	}{
		{
			inputId: validId,
			name:    "success",
			want: func(t *testing.T, app *cltest.TestApplication) *types.ChainStatus {
				return &types.ChainStatus{
					ID:      validId,
					Enabled: true,
					Config: `ChainID = 'Plugin-12'
BalancePollPeriod = '5s'
ConfirmPollPeriod = '500ms'
OCR2CachePollPeriod = '1s'
OCR2CacheTTL = '1m0s'
TxTimeout = '1h0m0s'
TxRetryTimeout = '10s'
TxConfirmTimeout = '30s'
SkipPreflight = false
Commitment = 'confirmed'
MaxRetries = 0
FeeEstimatorMode = 'fixed'
ComputeUnitPriceMax = 1000
ComputeUnitPriceMin = 0
ComputeUnitPriceDefault = 0
FeeBumpPeriod = '3s'
BlockHistoryPollPeriod = '5s'
Nodes = []
`,
				}
			},
			wantStatusCode: http.StatusOK,
		},
		{
			inputId: "234",
			name:    "not found",
			want: func(t *testing.T, app *cltest.TestApplication) *types.ChainStatus {
				return nil
			},
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			controller := setupSolanaChainsControllerTestV2(t, &solcfg.TOMLConfig{
				ChainID: ptr(validId),
				Chain: solcfg.Chain{
					SkipPreflight: ptr(false),
					TxTimeout:     commoncfg.MustNewDuration(time.Hour),
				},
			})

			wantedResult := tc.want(t, controller.app)
			resp, cleanup := controller.client.Get(
				fmt.Sprintf("/v2/chains/solana/%s", tc.inputId),
			)
			t.Cleanup(cleanup)
			require.Equal(t, tc.wantStatusCode, resp.StatusCode)

			if wantedResult != nil {
				resource1 := presenters.SolanaChainResource{}
				err := web.ParseJSONAPIResponse(cltest.ParseResponseBody(t, resp), &resource1)
				require.NoError(t, err)

				assert.Equal(t, wantedResult.ID, resource1.ID)
				assert.Equal(t, wantedResult.Enabled, resource1.Enabled)
				assert.Equal(t, wantedResult.Config, resource1.Config)
			}
		})
	}
}

func Test_SolanaChainsController_Index(t *testing.T) {
	t.Parallel()

	chainA := &solcfg.TOMLConfig{
		ChainID: ptr(fmt.Sprintf("PlugintestA-%d", rand.Int31n(999999))),
		Chain: solcfg.Chain{
			TxTimeout: commoncfg.MustNewDuration(time.Hour),
		},
	}

	chainB := &solcfg.TOMLConfig{
		ChainID: ptr(fmt.Sprintf("PlugintestB-%d", rand.Int31n(999999))),
		Chain: solcfg.Chain{
			SkipPreflight: ptr(false),
		},
	}
	controller := setupSolanaChainsControllerTestV2(t, chainA, chainB)

	badResp, cleanup := controller.client.Get("/v2/chains/solana?size=asd")
	t.Cleanup(cleanup)
	require.Equal(t, http.StatusUnprocessableEntity, badResp.StatusCode)

	resp, cleanup := controller.client.Get("/v2/chains/solana?size=1")
	t.Cleanup(cleanup)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body := cltest.ParseResponseBody(t, resp)

	metaCount, err := cltest.ParseJSONAPIResponseMetaCount(body)
	require.NoError(t, err)
	require.Equal(t, 2, metaCount)

	var links jsonapi.Links

	chains := []presenters.SolanaChainResource{}
	err = web.ParsePaginatedResponse(body, &chains, &links)
	assert.NoError(t, err)
	assert.NotEmpty(t, links["next"].Href)
	assert.Empty(t, links["prev"].Href)

	assert.Len(t, links, 1)
	assert.Equal(t, *chainA.ChainID, chains[0].ID)
	tomlA, err := chainA.TOMLString()
	require.NoError(t, err)
	assert.Equal(t, tomlA, chains[0].Config)

	resp, cleanup = controller.client.Get(links["next"].Href)
	t.Cleanup(cleanup)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	chains = []presenters.SolanaChainResource{}
	err = web.ParsePaginatedResponse(cltest.ParseResponseBody(t, resp), &chains, &links)
	assert.NoError(t, err)
	assert.Empty(t, links["next"].Href)
	assert.NotEmpty(t, links["prev"].Href)

	assert.Len(t, links, 1)
	assert.Equal(t, *chainB.ChainID, chains[0].ID)
	tomlB, err := chainB.TOMLString()
	require.NoError(t, err)
	assert.Equal(t, tomlB, chains[0].Config)
}

type TestSolanaChainsController struct {
	app    *cltest.TestApplication
	client cltest.HTTPClientCleaner
}

func setupSolanaChainsControllerTestV2(t *testing.T, cfgs ...*solcfg.TOMLConfig) *TestSolanaChainsController {
	for i := range cfgs {
		cfgs[i].SetDefaults()
	}
	cfg := configtest.NewGeneralConfig(t, func(c *plugin.Config, s *plugin.Secrets) {
		c.Solana = cfgs
		c.EVM = nil
	})
	app := cltest.NewApplicationWithConfig(t, cfg)
	require.NoError(t, app.Start(testutils.Context(t)))

	client := app.NewHTTPClient(nil)

	return &TestSolanaChainsController{
		app:    app,
		client: client,
	}
}
