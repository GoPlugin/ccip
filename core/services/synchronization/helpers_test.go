package synchronization

import (
	"net/url"
	"testing"
	"time"

	"github.com/goplugin/pluginv3.0/v2/core/logger"
	"github.com/goplugin/pluginv3.0/v2/core/services/keystore"
	telemPb "github.com/goplugin/pluginv3.0/v2/core/services/synchronization/telem"
)

// NewTestTelemetryIngressClient calls NewTelemetryIngressClient and injects telemClient.
func NewTestTelemetryIngressClient(t *testing.T, url *url.URL, serverPubKeyHex string, ks keystore.CSA, logging bool, telemClient telemPb.TelemClient) TelemetryService {
	tc := NewTelemetryIngressClient(url, serverPubKeyHex, ks, logging, logger.TestLogger(t), 100)
	tc.(*telemetryIngressClient).telemClient = telemClient
	return tc
}

// NewTestTelemetryIngressBatchClient calls NewTelemetryIngressBatchClient and injects telemClient.
func NewTestTelemetryIngressBatchClient(t *testing.T, url *url.URL, serverPubKeyHex string, ks keystore.CSA, logging bool, telemClient telemPb.TelemClient, sendInterval time.Duration, uniconn bool) TelemetryService {
	tc := NewTelemetryIngressBatchClient(url, serverPubKeyHex, ks, logging, logger.TestLogger(t), 100, 50, sendInterval, time.Second, uniconn)
	tc.(*telemetryIngressBatchClient).closeFn = func() error { return nil }
	tc.(*telemetryIngressBatchClient).telemClient = telemClient
	return tc
}
