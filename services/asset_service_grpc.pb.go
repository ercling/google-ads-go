// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: google/ads/googleads/v11/services/asset_service.proto

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

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Creates assets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AssetError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CurrencyCodeError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MediaUploadError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	//   [YoutubeVideoRegistrationError]()
	MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error) {
	out := new(MutateAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AssetService/MutateAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations must embed UnimplementedAssetServiceServer
// for forward compatibility
type AssetServiceServer interface {
	// Creates assets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AssetError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CurrencyCodeError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [ListOperationError]()
	//   [MediaUploadError]()
	//   [MutateError]()
	//   [NotAllowlistedError]()
	//   [NotEmptyError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	//   [YoutubeVideoRegistrationError]()
	MutateAssets(context.Context, *MutateAssetsRequest) (*MutateAssetsResponse, error)
	mustEmbedUnimplementedAssetServiceServer()
}

// UnimplementedAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (UnimplementedAssetServiceServer) MutateAssets(context.Context, *MutateAssetsRequest) (*MutateAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssets not implemented")
}
func (UnimplementedAssetServiceServer) mustEmbedUnimplementedAssetServiceServer() {}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_MutateAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).MutateAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.AssetService/MutateAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).MutateAssets(ctx, req.(*MutateAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssets",
			Handler:    _AssetService_MutateAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/asset_service.proto",
}
