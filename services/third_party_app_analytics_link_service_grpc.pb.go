// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: google/ads/googleads/v11/services/third_party_app_analytics_link_service.proto

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

// ThirdPartyAppAnalyticsLinkServiceClient is the client API for ThirdPartyAppAnalyticsLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdPartyAppAnalyticsLinkServiceClient interface {
	// Regenerate ThirdPartyAppAnalyticsLink.shareable_link_id that should be
	// provided to the third party when setting up app analytics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	RegenerateShareableLinkId(ctx context.Context, in *RegenerateShareableLinkIdRequest, opts ...grpc.CallOption) (*RegenerateShareableLinkIdResponse, error)
}

type thirdPartyAppAnalyticsLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyAppAnalyticsLinkServiceClient(cc grpc.ClientConnInterface) ThirdPartyAppAnalyticsLinkServiceClient {
	return &thirdPartyAppAnalyticsLinkServiceClient{cc}
}

func (c *thirdPartyAppAnalyticsLinkServiceClient) RegenerateShareableLinkId(ctx context.Context, in *RegenerateShareableLinkIdRequest, opts ...grpc.CallOption) (*RegenerateShareableLinkIdResponse, error) {
	out := new(RegenerateShareableLinkIdResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.ThirdPartyAppAnalyticsLinkService/RegenerateShareableLinkId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyAppAnalyticsLinkServiceServer is the server API for ThirdPartyAppAnalyticsLinkService service.
// All implementations must embed UnimplementedThirdPartyAppAnalyticsLinkServiceServer
// for forward compatibility
type ThirdPartyAppAnalyticsLinkServiceServer interface {
	// Regenerate ThirdPartyAppAnalyticsLink.shareable_link_id that should be
	// provided to the third party when setting up app analytics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	RegenerateShareableLinkId(context.Context, *RegenerateShareableLinkIdRequest) (*RegenerateShareableLinkIdResponse, error)
	mustEmbedUnimplementedThirdPartyAppAnalyticsLinkServiceServer()
}

// UnimplementedThirdPartyAppAnalyticsLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThirdPartyAppAnalyticsLinkServiceServer struct {
}

func (UnimplementedThirdPartyAppAnalyticsLinkServiceServer) RegenerateShareableLinkId(context.Context, *RegenerateShareableLinkIdRequest) (*RegenerateShareableLinkIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateShareableLinkId not implemented")
}
func (UnimplementedThirdPartyAppAnalyticsLinkServiceServer) mustEmbedUnimplementedThirdPartyAppAnalyticsLinkServiceServer() {
}

// UnsafeThirdPartyAppAnalyticsLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdPartyAppAnalyticsLinkServiceServer will
// result in compilation errors.
type UnsafeThirdPartyAppAnalyticsLinkServiceServer interface {
	mustEmbedUnimplementedThirdPartyAppAnalyticsLinkServiceServer()
}

func RegisterThirdPartyAppAnalyticsLinkServiceServer(s grpc.ServiceRegistrar, srv ThirdPartyAppAnalyticsLinkServiceServer) {
	s.RegisterService(&ThirdPartyAppAnalyticsLinkService_ServiceDesc, srv)
}

func _ThirdPartyAppAnalyticsLinkService_RegenerateShareableLinkId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateShareableLinkIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyAppAnalyticsLinkServiceServer).RegenerateShareableLinkId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.ThirdPartyAppAnalyticsLinkService/RegenerateShareableLinkId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyAppAnalyticsLinkServiceServer).RegenerateShareableLinkId(ctx, req.(*RegenerateShareableLinkIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdPartyAppAnalyticsLinkService_ServiceDesc is the grpc.ServiceDesc for ThirdPartyAppAnalyticsLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdPartyAppAnalyticsLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.ThirdPartyAppAnalyticsLinkService",
	HandlerType: (*ThirdPartyAppAnalyticsLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegenerateShareableLinkId",
			Handler:    _ThirdPartyAppAnalyticsLinkService_RegenerateShareableLinkId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/third_party_app_analytics_link_service.proto",
}
