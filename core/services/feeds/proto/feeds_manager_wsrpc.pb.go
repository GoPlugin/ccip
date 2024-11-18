// Code generated by protoc-gen-go-wsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-wsrpc v0.0.1
// - protoc             v4.25.3

package proto

import (
	context "context"

	wsrpc "github.com/goplugin/wsrpc"
)

// FeedsManagerClient is the client API for FeedsManager service.
type FeedsManagerClient interface {
	ApprovedJob(ctx context.Context, in *ApprovedJobRequest) (*ApprovedJobResponse, error)
	Healthcheck(ctx context.Context, in *HealthcheckRequest) (*HealthcheckResponse, error)
	UpdateNode(ctx context.Context, in *UpdateNodeRequest) (*UpdateNodeResponse, error)
	RejectedJob(ctx context.Context, in *RejectedJobRequest) (*RejectedJobResponse, error)
	CancelledJob(ctx context.Context, in *CancelledJobRequest) (*CancelledJobResponse, error)
}

type feedsManagerClient struct {
	cc wsrpc.ClientInterface
}

func NewFeedsManagerClient(cc wsrpc.ClientInterface) FeedsManagerClient {
	return &feedsManagerClient{cc}
}

func (c *feedsManagerClient) ApprovedJob(ctx context.Context, in *ApprovedJobRequest) (*ApprovedJobResponse, error) {
	out := new(ApprovedJobResponse)
	err := c.cc.Invoke(ctx, "ApprovedJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsManagerClient) Healthcheck(ctx context.Context, in *HealthcheckRequest) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "Healthcheck", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsManagerClient) UpdateNode(ctx context.Context, in *UpdateNodeRequest) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := c.cc.Invoke(ctx, "UpdateNode", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsManagerClient) RejectedJob(ctx context.Context, in *RejectedJobRequest) (*RejectedJobResponse, error) {
	out := new(RejectedJobResponse)
	err := c.cc.Invoke(ctx, "RejectedJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedsManagerClient) CancelledJob(ctx context.Context, in *CancelledJobRequest) (*CancelledJobResponse, error) {
	out := new(CancelledJobResponse)
	err := c.cc.Invoke(ctx, "CancelledJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedsManagerServer is the server API for FeedsManager service.
type FeedsManagerServer interface {
	ApprovedJob(context.Context, *ApprovedJobRequest) (*ApprovedJobResponse, error)
	Healthcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	UpdateNode(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
	RejectedJob(context.Context, *RejectedJobRequest) (*RejectedJobResponse, error)
	CancelledJob(context.Context, *CancelledJobRequest) (*CancelledJobResponse, error)
}

func RegisterFeedsManagerServer(s wsrpc.ServiceRegistrar, srv FeedsManagerServer) {
	s.RegisterService(&FeedsManager_ServiceDesc, srv)
}

func _FeedsManager_ApprovedJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ApprovedJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).ApprovedJob(ctx, in)
}

func _FeedsManager_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).Healthcheck(ctx, in)
}

func _FeedsManager_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).UpdateNode(ctx, in)
}

func _FeedsManager_RejectedJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RejectedJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).RejectedJob(ctx, in)
}

func _FeedsManager_CancelledJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CancelledJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(FeedsManagerServer).CancelledJob(ctx, in)
}

// FeedsManager_ServiceDesc is the wsrpc.ServiceDesc for FeedsManager service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedsManager_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "cfm.FeedsManager",
	HandlerType: (*FeedsManagerServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "ApprovedJob",
			Handler:    _FeedsManager_ApprovedJob_Handler,
		},
		{
			MethodName: "Healthcheck",
			Handler:    _FeedsManager_Healthcheck_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _FeedsManager_UpdateNode_Handler,
		},
		{
			MethodName: "RejectedJob",
			Handler:    _FeedsManager_RejectedJob_Handler,
		},
		{
			MethodName: "CancelledJob",
			Handler:    _FeedsManager_CancelledJob_Handler,
		},
	},
}

// NodeServiceClient is the client API for NodeService service.
type NodeServiceClient interface {
	ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest) (*DeleteJobResponse, error)
	RevokeJob(ctx context.Context, in *RevokeJobRequest) (*RevokeJobResponse, error)
}

type nodeServiceClient struct {
	cc wsrpc.ClientInterface
}

func NewNodeServiceClient(cc wsrpc.ClientInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) ProposeJob(ctx context.Context, in *ProposeJobRequest) (*ProposeJobResponse, error) {
	out := new(ProposeJobResponse)
	err := c.cc.Invoke(ctx, "ProposeJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) DeleteJob(ctx context.Context, in *DeleteJobRequest) (*DeleteJobResponse, error) {
	out := new(DeleteJobResponse)
	err := c.cc.Invoke(ctx, "DeleteJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) RevokeJob(ctx context.Context, in *RevokeJobRequest) (*RevokeJobResponse, error) {
	out := new(RevokeJobResponse)
	err := c.cc.Invoke(ctx, "RevokeJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	ProposeJob(context.Context, *ProposeJobRequest) (*ProposeJobResponse, error)
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error)
	RevokeJob(context.Context, *RevokeJobRequest) (*RevokeJobResponse, error)
}

func RegisterNodeServiceServer(s wsrpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_ProposeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProposeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).ProposeJob(ctx, in)
}

func _NodeService_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).DeleteJob(ctx, in)
}

func _NodeService_RevokeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RevokeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(NodeServiceServer).RevokeJob(ctx, in)
}

// NodeService_ServiceDesc is the wsrpc.ServiceDesc for NodeService service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "cfm.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "ProposeJob",
			Handler:    _NodeService_ProposeJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _NodeService_DeleteJob_Handler,
		},
		{
			MethodName: "RevokeJob",
			Handler:    _NodeService_RevokeJob_Handler,
		},
	},
}
