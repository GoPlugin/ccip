// Code generated by protoc-gen-go-wsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-wsrpc v0.0.1
// - protoc             v4.25.1

package pb

import (
	context "context"
	wsrpc "github.com/goplugin/wsrpc"
)

// MercuryClient is the client API for Mercury service.
type MercuryClient interface {
	Transmit(ctx context.Context, in *TransmitRequest) (*TransmitResponse, error)
	LatestReport(ctx context.Context, in *LatestReportRequest) (*LatestReportResponse, error)
}

type mercuryClient struct {
	cc wsrpc.ClientInterface
}

func NewMercuryClient(cc wsrpc.ClientInterface) MercuryClient {
	return &mercuryClient{cc}
}

func (c *mercuryClient) Transmit(ctx context.Context, in *TransmitRequest) (*TransmitResponse, error) {
	out := new(TransmitResponse)
	err := c.cc.Invoke(ctx, "Transmit", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryClient) LatestReport(ctx context.Context, in *LatestReportRequest) (*LatestReportResponse, error) {
	out := new(LatestReportResponse)
	err := c.cc.Invoke(ctx, "LatestReport", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryServer is the server API for Mercury service.
type MercuryServer interface {
	Transmit(context.Context, *TransmitRequest) (*TransmitResponse, error)
	LatestReport(context.Context, *LatestReportRequest) (*LatestReportResponse, error)
}

func RegisterMercuryServer(s wsrpc.ServiceRegistrar, srv MercuryServer) {
	s.RegisterService(&Mercury_ServiceDesc, srv)
}

func _Mercury_Transmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TransmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(MercuryServer).Transmit(ctx, in)
}

func _Mercury_LatestReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LatestReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(MercuryServer).LatestReport(ctx, in)
}

// Mercury_ServiceDesc is the wsrpc.ServiceDesc for Mercury service.
// It's only intended for direct use with wsrpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mercury_ServiceDesc = wsrpc.ServiceDesc{
	ServiceName: "pb.Mercury",
	HandlerType: (*MercuryServer)(nil),
	Methods: []wsrpc.MethodDesc{
		{
			MethodName: "Transmit",
			Handler:    _Mercury_Transmit_Handler,
		},
		{
			MethodName: "LatestReport",
			Handler:    _Mercury_LatestReport_Handler,
		},
	},
}
