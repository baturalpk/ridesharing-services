// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: mapping.proto

package genproto

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

// MappingServiceClient is the client API for MappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MappingServiceClient interface {
	CalculateDistance(ctx context.Context, in *CalculateDistanceRequest, opts ...grpc.CallOption) (*CalculateDistanceResponse, error)
	MakeTimeEstimation(ctx context.Context, in *MakeTimeEstimationRequest, opts ...grpc.CallOption) (*MakeTimeEstimationResponse, error)
	DoGeocoding(ctx context.Context, in *DoGeocodingRequest, opts ...grpc.CallOption) (*DoGeocodingResponse, error)
}

type mappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMappingServiceClient(cc grpc.ClientConnInterface) MappingServiceClient {
	return &mappingServiceClient{cc}
}

func (c *mappingServiceClient) CalculateDistance(ctx context.Context, in *CalculateDistanceRequest, opts ...grpc.CallOption) (*CalculateDistanceResponse, error) {
	out := new(CalculateDistanceResponse)
	err := c.cc.Invoke(ctx, "/mapping.MappingService/CalculateDistance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mappingServiceClient) MakeTimeEstimation(ctx context.Context, in *MakeTimeEstimationRequest, opts ...grpc.CallOption) (*MakeTimeEstimationResponse, error) {
	out := new(MakeTimeEstimationResponse)
	err := c.cc.Invoke(ctx, "/mapping.MappingService/MakeTimeEstimation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mappingServiceClient) DoGeocoding(ctx context.Context, in *DoGeocodingRequest, opts ...grpc.CallOption) (*DoGeocodingResponse, error) {
	out := new(DoGeocodingResponse)
	err := c.cc.Invoke(ctx, "/mapping.MappingService/DoGeocoding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MappingServiceServer is the server API for MappingService service.
// All implementations must embed UnimplementedMappingServiceServer
// for forward compatibility
type MappingServiceServer interface {
	CalculateDistance(context.Context, *CalculateDistanceRequest) (*CalculateDistanceResponse, error)
	MakeTimeEstimation(context.Context, *MakeTimeEstimationRequest) (*MakeTimeEstimationResponse, error)
	DoGeocoding(context.Context, *DoGeocodingRequest) (*DoGeocodingResponse, error)
	mustEmbedUnimplementedMappingServiceServer()
}

// UnimplementedMappingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMappingServiceServer struct {
}

func (UnimplementedMappingServiceServer) CalculateDistance(context.Context, *CalculateDistanceRequest) (*CalculateDistanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateDistance not implemented")
}
func (UnimplementedMappingServiceServer) MakeTimeEstimation(context.Context, *MakeTimeEstimationRequest) (*MakeTimeEstimationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeTimeEstimation not implemented")
}
func (UnimplementedMappingServiceServer) DoGeocoding(context.Context, *DoGeocodingRequest) (*DoGeocodingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoGeocoding not implemented")
}
func (UnimplementedMappingServiceServer) mustEmbedUnimplementedMappingServiceServer() {}

// UnsafeMappingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MappingServiceServer will
// result in compilation errors.
type UnsafeMappingServiceServer interface {
	mustEmbedUnimplementedMappingServiceServer()
}

func RegisterMappingServiceServer(s grpc.ServiceRegistrar, srv MappingServiceServer) {
	s.RegisterService(&MappingService_ServiceDesc, srv)
}

func _MappingService_CalculateDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateDistanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).CalculateDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapping.MappingService/CalculateDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).CalculateDistance(ctx, req.(*CalculateDistanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MappingService_MakeTimeEstimation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeTimeEstimationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).MakeTimeEstimation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapping.MappingService/MakeTimeEstimation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).MakeTimeEstimation(ctx, req.(*MakeTimeEstimationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MappingService_DoGeocoding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoGeocodingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).DoGeocoding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapping.MappingService/DoGeocoding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).DoGeocoding(ctx, req.(*DoGeocodingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MappingService_ServiceDesc is the grpc.ServiceDesc for MappingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MappingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mapping.MappingService",
	HandlerType: (*MappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateDistance",
			Handler:    _MappingService_CalculateDistance_Handler,
		},
		{
			MethodName: "MakeTimeEstimation",
			Handler:    _MappingService_MakeTimeEstimation_Handler,
		},
		{
			MethodName: "DoGeocoding",
			Handler:    _MappingService_DoGeocoding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapping.proto",
}
