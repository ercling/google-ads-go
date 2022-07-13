// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: google/ads/googleads/v11/services/campaign_shared_set_service.proto

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

// CampaignSharedSetServiceClient is the client API for CampaignSharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignSharedSetServiceClient interface {
	// Creates or removes campaign shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CampaignSharedSetError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCampaignSharedSets(ctx context.Context, in *MutateCampaignSharedSetsRequest, opts ...grpc.CallOption) (*MutateCampaignSharedSetsResponse, error)
}

type campaignSharedSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignSharedSetServiceClient(cc grpc.ClientConnInterface) CampaignSharedSetServiceClient {
	return &campaignSharedSetServiceClient{cc}
}

func (c *campaignSharedSetServiceClient) MutateCampaignSharedSets(ctx context.Context, in *MutateCampaignSharedSetsRequest, opts ...grpc.CallOption) (*MutateCampaignSharedSetsResponse, error) {
	out := new(MutateCampaignSharedSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.CampaignSharedSetService/MutateCampaignSharedSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignSharedSetServiceServer is the server API for CampaignSharedSetService service.
// All implementations must embed UnimplementedCampaignSharedSetServiceServer
// for forward compatibility
type CampaignSharedSetServiceServer interface {
	// Creates or removes campaign shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CampaignSharedSetError]()
	//   [ContextError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateCampaignSharedSets(context.Context, *MutateCampaignSharedSetsRequest) (*MutateCampaignSharedSetsResponse, error)
	mustEmbedUnimplementedCampaignSharedSetServiceServer()
}

// UnimplementedCampaignSharedSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignSharedSetServiceServer struct {
}

func (UnimplementedCampaignSharedSetServiceServer) MutateCampaignSharedSets(context.Context, *MutateCampaignSharedSetsRequest) (*MutateCampaignSharedSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignSharedSets not implemented")
}
func (UnimplementedCampaignSharedSetServiceServer) mustEmbedUnimplementedCampaignSharedSetServiceServer() {
}

// UnsafeCampaignSharedSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignSharedSetServiceServer will
// result in compilation errors.
type UnsafeCampaignSharedSetServiceServer interface {
	mustEmbedUnimplementedCampaignSharedSetServiceServer()
}

func RegisterCampaignSharedSetServiceServer(s grpc.ServiceRegistrar, srv CampaignSharedSetServiceServer) {
	s.RegisterService(&CampaignSharedSetService_ServiceDesc, srv)
}

func _CampaignSharedSetService_MutateCampaignSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignSharedSetServiceServer).MutateCampaignSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v11.services.CampaignSharedSetService/MutateCampaignSharedSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignSharedSetServiceServer).MutateCampaignSharedSets(ctx, req.(*MutateCampaignSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignSharedSetService_ServiceDesc is the grpc.ServiceDesc for CampaignSharedSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignSharedSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.CampaignSharedSetService",
	HandlerType: (*CampaignSharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignSharedSets",
			Handler:    _CampaignSharedSetService_MutateCampaignSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/campaign_shared_set_service.proto",
}
