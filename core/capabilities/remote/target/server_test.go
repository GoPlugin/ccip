package target_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commoncap "github.com/goplugin/plugin-common/pkg/capabilities"
	"github.com/goplugin/plugin-common/pkg/capabilities/pb"
	"github.com/goplugin/plugin-common/pkg/services"
	"github.com/goplugin/plugin-common/pkg/values"
	"github.com/goplugin/pluginv3.0/v2/core/capabilities/remote/target"
	remotetypes "github.com/goplugin/pluginv3.0/v2/core/capabilities/remote/types"
	"github.com/goplugin/pluginv3.0/v2/core/internal/testutils"
	"github.com/goplugin/pluginv3.0/v2/core/logger"
	p2ptypes "github.com/goplugin/pluginv3.0/v2/core/services/p2p/types"
)

func Test_Server_ExcludesNonDeterministicInputAttributes(t *testing.T) {
	ctx := testutils.Context(t)

	numCapabilityPeers := 4

	callers, srvcs := testRemoteTargetServer(ctx, t, &commoncap.RemoteTargetConfig{RequestHashExcludedAttributes: []string{"signed_report.Signatures"}},
		&TestCapability{}, 10, 9, numCapabilityPeers, 3, 10*time.Minute)

	for idx, caller := range callers {
		rawInputs := map[string]any{
			"signed_report": map[string]any{"Signatures": "sig" + strconv.Itoa(idx), "Price": 20},
		}

		inputs, err := values.NewMap(rawInputs)
		require.NoError(t, err)

		_, err = caller.Execute(context.Background(),
			commoncap.CapabilityRequest{
				Metadata: commoncap.RequestMetadata{
					WorkflowID:          workflowID1,
					WorkflowExecutionID: workflowExecutionID1,
				},
				Inputs: inputs,
			})
		require.NoError(t, err)
	}

	for _, caller := range callers {
		for i := 0; i < numCapabilityPeers; i++ {
			msg := <-caller.receivedMessages
			assert.Equal(t, remotetypes.Error_OK, msg.Error)
		}
	}
	closeServices(t, srvcs)
}

func Test_Server_RespondsAfterSufficientRequests(t *testing.T) {
	ctx := testutils.Context(t)

	numCapabilityPeers := 4

	callers, srvcs := testRemoteTargetServer(ctx, t, &commoncap.RemoteTargetConfig{}, &TestCapability{}, 10, 9, numCapabilityPeers, 3, 10*time.Minute)

	for _, caller := range callers {
		_, err := caller.Execute(context.Background(),
			commoncap.CapabilityRequest{
				Metadata: commoncap.RequestMetadata{
					WorkflowID:          workflowID1,
					WorkflowExecutionID: workflowExecutionID1,
				},
			})
		require.NoError(t, err)
	}

	for _, caller := range callers {
		for i := 0; i < numCapabilityPeers; i++ {
			msg := <-caller.receivedMessages
			assert.Equal(t, remotetypes.Error_OK, msg.Error)
		}
	}
	closeServices(t, srvcs)
}

func Test_Server_InsufficientCallers(t *testing.T) {
	ctx := testutils.Context(t)

	numCapabilityPeers := 4

	callers, srvcs := testRemoteTargetServer(ctx, t, &commoncap.RemoteTargetConfig{}, &TestCapability{}, 10, 10, numCapabilityPeers, 3, 100*time.Millisecond)

	for _, caller := range callers {
		_, err := caller.Execute(context.Background(),
			commoncap.CapabilityRequest{
				Metadata: commoncap.RequestMetadata{
					WorkflowID:          workflowID1,
					WorkflowExecutionID: workflowExecutionID1,
				},
			})
		require.NoError(t, err)
	}

	for _, caller := range callers {
		for i := 0; i < numCapabilityPeers; i++ {
			msg := <-caller.receivedMessages
			assert.Equal(t, remotetypes.Error_TIMEOUT, msg.Error)
		}
	}
	closeServices(t, srvcs)
}

func Test_Server_CapabilityError(t *testing.T) {
	ctx := testutils.Context(t)

	numCapabilityPeers := 4

	callers, srvcs := testRemoteTargetServer(ctx, t, &commoncap.RemoteTargetConfig{}, &TestErrorCapability{}, 10, 9, numCapabilityPeers, 3, 100*time.Millisecond)

	for _, caller := range callers {
		_, err := caller.Execute(context.Background(),
			commoncap.CapabilityRequest{
				Metadata: commoncap.RequestMetadata{
					WorkflowID:          workflowID1,
					WorkflowExecutionID: workflowExecutionID1,
				},
			})
		require.NoError(t, err)
	}

	for _, caller := range callers {
		for i := 0; i < numCapabilityPeers; i++ {
			msg := <-caller.receivedMessages
			assert.Equal(t, remotetypes.Error_INTERNAL_ERROR, msg.Error)
		}
	}
	closeServices(t, srvcs)
}

func testRemoteTargetServer(ctx context.Context, t *testing.T,
	config *commoncap.RemoteTargetConfig,
	underlying commoncap.TargetCapability,
	numWorkflowPeers int, workflowDonF uint8,
	numCapabilityPeers int, capabilityDonF uint8, capabilityNodeResponseTimeout time.Duration) ([]*serverTestClient, []services.Service) {
	lggr := logger.TestLogger(t)

	capabilityPeers := make([]p2ptypes.PeerID, numCapabilityPeers)
	for i := 0; i < numCapabilityPeers; i++ {
		capabilityPeerID := NewP2PPeerID(t)
		capabilityPeers[i] = capabilityPeerID
	}

	capDonInfo := commoncap.DON{
		ID:      1,
		Members: capabilityPeers,
		F:       capabilityDonF,
	}

	capInfo := commoncap.CapabilityInfo{
		ID:             "cap_id@1.0.0",
		CapabilityType: commoncap.CapabilityTypeTarget,
		Description:    "Remote Target",
		DON:            &capDonInfo,
	}

	workflowPeers := make([]p2ptypes.PeerID, numWorkflowPeers)
	for i := 0; i < numWorkflowPeers; i++ {
		workflowPeers[i] = NewP2PPeerID(t)
	}

	workflowDonInfo := commoncap.DON{
		Members: workflowPeers,
		ID:      2,
		F:       workflowDonF,
	}

	var srvcs []services.Service
	broker := newTestAsyncMessageBroker(t, 1000)
	err := broker.Start(context.Background())
	require.NoError(t, err)
	srvcs = append(srvcs, broker)

	workflowDONs := map[uint32]commoncap.DON{
		workflowDonInfo.ID: workflowDonInfo,
	}

	capabilityNodes := make([]remotetypes.Receiver, numCapabilityPeers)

	for i := 0; i < numCapabilityPeers; i++ {
		capabilityPeer := capabilityPeers[i]
		capabilityDispatcher := broker.NewDispatcherForNode(capabilityPeer)
		capabilityNode := target.NewServer(config, capabilityPeer, underlying, capInfo, capDonInfo, workflowDONs, capabilityDispatcher,
			capabilityNodeResponseTimeout, lggr)
		require.NoError(t, capabilityNode.Start(ctx))
		broker.RegisterReceiverNode(capabilityPeer, capabilityNode)
		capabilityNodes[i] = capabilityNode
		srvcs = append(srvcs, capabilityNode)
	}

	workflowNodes := make([]*serverTestClient, numWorkflowPeers)
	for i := 0; i < numWorkflowPeers; i++ {
		workflowPeerDispatcher := broker.NewDispatcherForNode(workflowPeers[i])
		workflowNode := newServerTestClient(workflowPeers[i], capDonInfo, workflowPeerDispatcher)
		broker.RegisterReceiverNode(workflowPeers[i], workflowNode)
		workflowNodes[i] = workflowNode
	}

	return workflowNodes, srvcs
}

func closeServices(t *testing.T, srvcs []services.Service) {
	for _, srv := range srvcs {
		require.NoError(t, srv.Close())
	}
}

type serverTestClient struct {
	peerID            p2ptypes.PeerID
	dispatcher        remotetypes.Dispatcher
	capabilityDonInfo commoncap.DON
	receivedMessages  chan *remotetypes.MessageBody
	callerDonID       string
}

func (r *serverTestClient) Receive(_ context.Context, msg *remotetypes.MessageBody) {
	r.receivedMessages <- msg
}

func newServerTestClient(peerID p2ptypes.PeerID, capabilityDonInfo commoncap.DON,
	dispatcher remotetypes.Dispatcher) *serverTestClient {
	return &serverTestClient{peerID: peerID, dispatcher: dispatcher, capabilityDonInfo: capabilityDonInfo,
		receivedMessages: make(chan *remotetypes.MessageBody, 100), callerDonID: "workflow-don"}
}

func (r *serverTestClient) Info(ctx context.Context) (commoncap.CapabilityInfo, error) {
	panic("not implemented")
}

func (r *serverTestClient) RegisterToWorkflow(ctx context.Context, request commoncap.RegisterToWorkflowRequest) error {
	panic("not implemented")
}

func (r *serverTestClient) UnregisterFromWorkflow(ctx context.Context, request commoncap.UnregisterFromWorkflowRequest) error {
	panic("not implemented")
}

func (r *serverTestClient) Execute(ctx context.Context, req commoncap.CapabilityRequest) (<-chan commoncap.CapabilityResponse, error) {
	rawRequest, err := pb.MarshalCapabilityRequest(req)
	if err != nil {
		return nil, err
	}

	messageID, err := target.GetMessageIDForRequest(req)
	if err != nil {
		return nil, err
	}

	for _, node := range r.capabilityDonInfo.Members {
		message := &remotetypes.MessageBody{
			CapabilityId:    "capability-id",
			CapabilityDonId: 1,
			CallerDonId:     2,
			Method:          remotetypes.MethodExecute,
			Payload:         rawRequest,
			MessageId:       []byte(messageID),
			Sender:          r.peerID[:],
			Receiver:        node[:],
		}

		if err = r.dispatcher.Send(node, message); err != nil {
			return nil, err
		}
	}

	return nil, nil
}
