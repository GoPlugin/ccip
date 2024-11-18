package types

import (
	ctf_config "github.com/goplugin/plugin-testing-framework/lib/config"
	"github.com/goplugin/plugin-testing-framework/lib/testreporters"
	tc "github.com/goplugin/pluginv3.0/integration-tests/testconfig"
)

type VRFv2TestConfig interface {
	tc.CommonTestConfig
	ctf_config.GlobalTestConfig
	tc.VRFv2TestConfig
}

type VRFv2PlusTestConfig interface {
	tc.CommonTestConfig
	ctf_config.GlobalTestConfig
	tc.VRFv2PlusTestConfig
}

type FunctionsTestConfig interface {
	tc.CommonTestConfig
	ctf_config.GlobalTestConfig
	tc.FunctionsTestConfig
}

type AutomationTestConfig interface {
	ctf_config.GlobalTestConfig
	tc.CommonTestConfig
	tc.UpgradeablePluginTestConfig
	tc.AutomationTestConfig
}

type AutomationBenchmarkTestConfig interface {
	ctf_config.GlobalTestConfig
	tc.CommonTestConfig
	tc.AutomationTestConfig
	ctf_config.NamedConfigurations
	testreporters.GrafanaURLProvider
}

type OcrTestConfig interface {
	ctf_config.GlobalTestConfig
	tc.CommonTestConfig
	tc.OcrTestConfig
	ctf_config.SethConfig
}

type Ocr2TestConfig interface {
	ctf_config.GlobalTestConfig
	tc.CommonTestConfig
	tc.Ocr2TestConfig
}

type CCIPTestConfig interface {
	ctf_config.GlobalTestConfig
	tc.CommonTestConfig
	tc.CCIPTestConfig
}