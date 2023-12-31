// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: internal/riotgrpc/proto/summoner.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RiotSummonerClient is the client API for RiotSummoner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiotSummonerClient interface {
	ByName(ctx context.Context, in *ByNameRequest, opts ...grpc.CallOption) (*Summoner, error)
	ByPuuid(ctx context.Context, in *ByPuuidRequest, opts ...grpc.CallOption) (*Summoner, error)
	BySummonerId(ctx context.Context, in *BySummonerIdRequest, opts ...grpc.CallOption) (*Summoner, error)
	GetSoloq(ctx context.Context, in *BySummonerIdRequest, opts ...grpc.CallOption) (*LeagueEntry, error)
	GetMatchlist(ctx context.Context, in *ByPuuidMatchlistRequest, opts ...grpc.CallOption) (*Matchlist, error)
}

type riotSummonerClient struct {
	cc grpc.ClientConnInterface
}

func NewRiotSummonerClient(cc grpc.ClientConnInterface) RiotSummonerClient {
	return &riotSummonerClient{cc}
}

func (c *riotSummonerClient) ByName(ctx context.Context, in *ByNameRequest, opts ...grpc.CallOption) (*Summoner, error) {
	out := new(Summoner)
	err := c.cc.Invoke(ctx, "/RiotSummoner/ByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotSummonerClient) ByPuuid(ctx context.Context, in *ByPuuidRequest, opts ...grpc.CallOption) (*Summoner, error) {
	out := new(Summoner)
	err := c.cc.Invoke(ctx, "/RiotSummoner/ByPuuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotSummonerClient) BySummonerId(ctx context.Context, in *BySummonerIdRequest, opts ...grpc.CallOption) (*Summoner, error) {
	out := new(Summoner)
	err := c.cc.Invoke(ctx, "/RiotSummoner/BySummonerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotSummonerClient) GetSoloq(ctx context.Context, in *BySummonerIdRequest, opts ...grpc.CallOption) (*LeagueEntry, error) {
	out := new(LeagueEntry)
	err := c.cc.Invoke(ctx, "/RiotSummoner/GetSoloq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotSummonerClient) GetMatchlist(ctx context.Context, in *ByPuuidMatchlistRequest, opts ...grpc.CallOption) (*Matchlist, error) {
	out := new(Matchlist)
	err := c.cc.Invoke(ctx, "/RiotSummoner/GetMatchlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiotSummonerServer is the server API for RiotSummoner service.
// All implementations must embed UnimplementedRiotSummonerServer
// for forward compatibility
type RiotSummonerServer interface {
	ByName(context.Context, *ByNameRequest) (*Summoner, error)
	ByPuuid(context.Context, *ByPuuidRequest) (*Summoner, error)
	BySummonerId(context.Context, *BySummonerIdRequest) (*Summoner, error)
	GetSoloq(context.Context, *BySummonerIdRequest) (*LeagueEntry, error)
	GetMatchlist(context.Context, *ByPuuidMatchlistRequest) (*Matchlist, error)
	mustEmbedUnimplementedRiotSummonerServer()
}

// UnimplementedRiotSummonerServer must be embedded to have forward compatible implementations.
type UnimplementedRiotSummonerServer struct {
}

func (UnimplementedRiotSummonerServer) ByName(context.Context, *ByNameRequest) (*Summoner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByName not implemented")
}
func (UnimplementedRiotSummonerServer) ByPuuid(context.Context, *ByPuuidRequest) (*Summoner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByPuuid not implemented")
}
func (UnimplementedRiotSummonerServer) BySummonerId(context.Context, *BySummonerIdRequest) (*Summoner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BySummonerId not implemented")
}
func (UnimplementedRiotSummonerServer) GetSoloq(context.Context, *BySummonerIdRequest) (*LeagueEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSoloq not implemented")
}
func (UnimplementedRiotSummonerServer) GetMatchlist(context.Context, *ByPuuidMatchlistRequest) (*Matchlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchlist not implemented")
}
func (UnimplementedRiotSummonerServer) mustEmbedUnimplementedRiotSummonerServer() {}

// UnsafeRiotSummonerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RiotSummonerServer will
// result in compilation errors.
type UnsafeRiotSummonerServer interface {
	mustEmbedUnimplementedRiotSummonerServer()
}

func RegisterRiotSummonerServer(s grpc.ServiceRegistrar, srv RiotSummonerServer) {
	s.RegisterService(&RiotSummoner_ServiceDesc, srv)
}

func _RiotSummoner_ByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotSummonerServer).ByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RiotSummoner/ByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotSummonerServer).ByName(ctx, req.(*ByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotSummoner_ByPuuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByPuuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotSummonerServer).ByPuuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RiotSummoner/ByPuuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotSummonerServer).ByPuuid(ctx, req.(*ByPuuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotSummoner_BySummonerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BySummonerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotSummonerServer).BySummonerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RiotSummoner/BySummonerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotSummonerServer).BySummonerId(ctx, req.(*BySummonerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotSummoner_GetSoloq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BySummonerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotSummonerServer).GetSoloq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RiotSummoner/GetSoloq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotSummonerServer).GetSoloq(ctx, req.(*BySummonerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotSummoner_GetMatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByPuuidMatchlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotSummonerServer).GetMatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RiotSummoner/GetMatchlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotSummonerServer).GetMatchlist(ctx, req.(*ByPuuidMatchlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RiotSummoner_ServiceDesc is the grpc.ServiceDesc for RiotSummoner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RiotSummoner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RiotSummoner",
	HandlerType: (*RiotSummonerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByName",
			Handler:    _RiotSummoner_ByName_Handler,
		},
		{
			MethodName: "ByPuuid",
			Handler:    _RiotSummoner_ByPuuid_Handler,
		},
		{
			MethodName: "BySummonerId",
			Handler:    _RiotSummoner_BySummonerId_Handler,
		},
		{
			MethodName: "GetSoloq",
			Handler:    _RiotSummoner_GetSoloq_Handler,
		},
		{
			MethodName: "GetMatchlist",
			Handler:    _RiotSummoner_GetMatchlist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/riotgrpc/proto/summoner.proto",
}
