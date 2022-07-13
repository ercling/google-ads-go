// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: google/ads/googleads/v11/services/customer_asset_service.proto

package services

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

// CustomerAssetServiceClient is the client API for CustomerAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerAssetServiceClient interface {
	// Creates, updates, or removes customer assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerAssets(ctx context.Context, in *MutateCustomerAssetsRequest, opts ...grpc.CallOption) (*MutateCustomerAssetsResponse, error)
}

type customerAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerAssetServiceClient(cc grpc.ClientConnInterface) CustomerAssetServiceClient {
	return &customerAssetServiceClient{cc}
}

func (c *customerAssetServiceClient) MutateCustomerAssets(ctx context.Context, in *MutateCustomerAssetsRequest, opts ...grpc.CallOption) (*MutateCustomerAssetsResponse, error) {
	out := new(MutateCustomerAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.CustomerAssetService/MutateCustomerAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerAssetServiceServer is the server API for CustomerAssetService service.
// All implementations must embed UnimplementedCustomerAssetServiceServer
// for forward compatibility
type CustomerAssetServiceServer interface {
	// Creates, updates, or removes customer assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AssetLinkError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerAssets(context.Context, *MutateCustomerAssetsRequest) (*MutateCustomerAssetsResponse, error)
	mustEmbedUnimplementedCustomerAssetServiceServer()
}

// UnimplementedCustomerAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerAssetServiceServer struct {
}

func (UnimplementedCustomerAssetServiceServer) MutateCustomerAssets(context.Context, *MutateCustomerAssetsRequest) (*MutateCustomerAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerAssets not implemented")
}
func (UnimplementedCustomerAssetServiceServer) mustEmbedUnimplementedCustomerAssetServiceServer() {}

// UnsafeCustomerAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerAssetServiceServer will
// result in compilation errors.
type UnsafeCustomerAssetServiceServer interface {
	mustEmbedUnimplementedCustomerAssetServiceServer()
}

func RegisterCustomerAssetServiceServer(s grpc.ServiceRegistrar, srv CustomerAssetServiceServer) {
	s.RegisterService(&CustomerAssetService_ServiceDesc, srv)
}

func _CustomerAssetService_MutateCustomerAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAssetServiceServer).MutateCustomerAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.CustomerAssetService/MutateCustomerAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAssetServiceServer).MutateCustomerAssets(ctx, req.(*MutateCustomerAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerAssetService_ServiceDesc is the grpc.ServiceDesc for CustomerAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.CustomerAssetService",
	HandlerType: (*CustomerAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerAssets",
			Handler:    _CustomerAssetService_MutateCustomerAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/customer_asset_service.proto",
}
