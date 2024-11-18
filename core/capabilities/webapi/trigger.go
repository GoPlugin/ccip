package webapi

import (
	"github.com/goplugin/plugin-common/pkg/types/core"
	"github.com/goplugin/pluginv3.0/v2/core/logger"
	"github.com/goplugin/pluginv3.0/v2/core/services/gateway/connector"
	"github.com/goplugin/pluginv3.0/v2/core/services/job"
)

func NewTrigger(config string, registry core.CapabilitiesRegistry, connector connector.GatewayConnector, lggr logger.Logger) (job.ServiceCtx, error) {
	// TODO (CAPPL-22, CAPPL-24):
	//   - decode config
	//   - create an implementation of the capability API and add it to the Registry
	//   - create a handler and register it with Gateway Connector
	//   - manage trigger subscriptions
	//   - process incoming trigger events and related metadata
	return nil, nil
}
