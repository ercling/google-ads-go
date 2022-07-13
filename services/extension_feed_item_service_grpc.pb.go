// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: google/ads/googleads/v11/services/extension_feed_item_service.proto

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

// ExtensionFeedItemServiceClient is the client API for ExtensionFeedItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtensionFeedItemServiceClient interface {
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CountryCodeError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [ExtensionFeedItemError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [LanguageCodeError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [OperationAccessDeniedError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error)
}

type extensionFeedItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExtensionFeedItemServiceClient(cc grpc.ClientConnInterface) ExtensionFeedItemServiceClient {
	return &extensionFeedItemServiceClient{cc}
}

func (c *extensionFeedItemServiceClient) MutateExtensionFeedItems(ctx context.Context, in *MutateExtensionFeedItemsRequest, opts ...grpc.CallOption) (*MutateExtensionFeedItemsResponse, error) {
	out := new(MutateExtensionFeedItemsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.ExtensionFeedItemService/MutateExtensionFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionFeedItemServiceServer is the server API for ExtensionFeedItemService service.
// All implementations must embed UnimplementedExtensionFeedItemServiceServer
// for forward compatibility
type ExtensionFeedItemServiceServer interface {
	// Creates, updates, or removes extension feed items. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CollectionSizeError]()
	//   [CountryCodeError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [ExtensionFeedItemError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [ImageError]()
	//   [InternalError]()
	//   [LanguageCodeError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [OperationAccessDeniedError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringLengthError]()
	//   [UrlFieldError]()
	MutateExtensionFeedItems(context.Context, *MutateExtensionFeedItemsRequest) (*MutateExtensionFeedItemsResponse, error)
	mustEmbedUnimplementedExtensionFeedItemServiceServer()
}

// UnimplementedExtensionFeedItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExtensionFeedItemServiceServer struct {
}

func (UnimplementedExtensionFeedItemServiceServer) MutateExtensionFeedItems(context.Context, *MutateExtensionFeedItemsRequest) (*MutateExtensionFeedItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateExtensionFeedItems not implemented")
}
func (UnimplementedExtensionFeedItemServiceServer) mustEmbedUnimplementedExtensionFeedItemServiceServer() {
}

// UnsafeExtensionFeedItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtensionFeedItemServiceServer will
// result in compilation errors.
type UnsafeExtensionFeedItemServiceServer interface {
	mustEmbedUnimplementedExtensionFeedItemServiceServer()
}

func RegisterExtensionFeedItemServiceServer(s grpc.ServiceRegistrar, srv ExtensionFeedItemServiceServer) {
	s.RegisterService(&ExtensionFeedItemService_ServiceDesc, srv)
}

func _ExtensionFeedItemService_MutateExtensionFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateExtensionFeedItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.ExtensionFeedItemService/MutateExtensionFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtensionFeedItemServiceServer).MutateExtensionFeedItems(ctx, req.(*MutateExtensionFeedItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExtensionFeedItemService_ServiceDesc is the grpc.ServiceDesc for ExtensionFeedItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExtensionFeedItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.ExtensionFeedItemService",
	HandlerType: (*ExtensionFeedItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateExtensionFeedItems",
			Handler:    _ExtensionFeedItemService_MutateExtensionFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/extension_feed_item_service.proto",
}
