// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pb/Job/job.proto

package Job

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

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	PostJob(ctx context.Context, in *PostjobReq, opts ...grpc.CallOption) (*PostjobRes, error)
	EditJob(ctx context.Context, in *EditjobReq, opts ...grpc.CallOption) (*EditjobRes, error)
	SendProposal(ctx context.Context, in *ProposalReq, opts ...grpc.CallOption) (*ProposalRes, error)
	SendOffer(ctx context.Context, in *SendOfferReq, opts ...grpc.CallOption) (*SendOfferRes, error)
	AcceptOffer(ctx context.Context, in *AcceptOfferReq, opts ...grpc.CallOption) (*AcceptOfferRes, error)
	ViewContract(ctx context.Context, in *ContractReq, opts ...grpc.CallOption) (*ViewContractRes, error)
	AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryRes, error)
	GetCategory(ctx context.Context, in *GetCategoryReq, opts ...grpc.CallOption) (*GetCategoryRes, error)
	GetMyJobs(ctx context.Context, in *GetMyJobsReq, opts ...grpc.CallOption) (*GetMyJobsRes, error)
	GetJob(ctx context.Context, in *GetJobReq, opts ...grpc.CallOption) (*GetJobRes, error)
	GetJobs(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*GetJobsRes, error)
	SendWeeklyInvoice(ctx context.Context, in *InvoiceReq, opts ...grpc.CallOption) (*InvoiceRes, error)
	GetJobProposals(ctx context.Context, in *GJPReq, opts ...grpc.CallOption) (*GJPRes, error)
	GetJobOffersForFreelancer(ctx context.Context, in *GetJobOfferForFreelancerReq, opts ...grpc.CallOption) (*GetJobOfferForFreelancerRes, error)
	SearchJobs(ctx context.Context, in *SearchJobsReq, opts ...grpc.CallOption) (*SearchJobsRes, error)
	GetAllContractsForClient(ctx context.Context, in *GetAllContractsForClientReq, opts ...grpc.CallOption) (*GetAllContractsForClientRes, error)
	GetOneContractForClient(ctx context.Context, in *GetOneContractForClientReq, opts ...grpc.CallOption) (*GetOneContractForClientRes, error)
	ExecutePaymentContract(ctx context.Context, in *ExecutePaymentReq, opts ...grpc.CallOption) (*ExecutePaymentRes, error)
	CapturePaymentContract(ctx context.Context, in *CapturePaymentReq, opts ...grpc.CallOption) (*CapturePaymentRes, error)
	SaveJobs(ctx context.Context, in *SaveJobsReq, opts ...grpc.CallOption) (*SaveJobsRes, error)
	GetInvoiceContract(ctx context.Context, in *GetInvoiceContractReq, opts ...grpc.CallOption) (*GetInvoiceContractRes, error)
	AddAttachmentToContract(ctx context.Context, in *AddAttachmentReq, opts ...grpc.CallOption) (*AddAttachmentRes, error)
	GetAttachments(ctx context.Context, in *GetAttachmentReq, opts ...grpc.CallOption) (*GetAttachmentRes, error)
	CheckInvoiceStatus(ctx context.Context, in *CheckInvoiceStatusReq, opts ...grpc.CallOption) (*CheckInvoiceStatusRes, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) PostJob(ctx context.Context, in *PostjobReq, opts ...grpc.CallOption) (*PostjobRes, error) {
	out := new(PostjobRes)
	err := c.cc.Invoke(ctx, "/job.JobService/PostJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) EditJob(ctx context.Context, in *EditjobReq, opts ...grpc.CallOption) (*EditjobRes, error) {
	out := new(EditjobRes)
	err := c.cc.Invoke(ctx, "/job.JobService/EditJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SendProposal(ctx context.Context, in *ProposalReq, opts ...grpc.CallOption) (*ProposalRes, error) {
	out := new(ProposalRes)
	err := c.cc.Invoke(ctx, "/job.JobService/SendProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SendOffer(ctx context.Context, in *SendOfferReq, opts ...grpc.CallOption) (*SendOfferRes, error) {
	out := new(SendOfferRes)
	err := c.cc.Invoke(ctx, "/job.JobService/SendOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) AcceptOffer(ctx context.Context, in *AcceptOfferReq, opts ...grpc.CallOption) (*AcceptOfferRes, error) {
	out := new(AcceptOfferRes)
	err := c.cc.Invoke(ctx, "/job.JobService/AcceptOffer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ViewContract(ctx context.Context, in *ContractReq, opts ...grpc.CallOption) (*ViewContractRes, error) {
	out := new(ViewContractRes)
	err := c.cc.Invoke(ctx, "/job.JobService/ViewContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryRes, error) {
	out := new(AddCategoryRes)
	err := c.cc.Invoke(ctx, "/job.JobService/AddCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetCategory(ctx context.Context, in *GetCategoryReq, opts ...grpc.CallOption) (*GetCategoryRes, error) {
	out := new(GetCategoryRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetMyJobs(ctx context.Context, in *GetMyJobsReq, opts ...grpc.CallOption) (*GetMyJobsRes, error) {
	out := new(GetMyJobsRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetMyJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *GetJobReq, opts ...grpc.CallOption) (*GetJobRes, error) {
	out := new(GetJobRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobs(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*GetJobsRes, error) {
	out := new(GetJobsRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SendWeeklyInvoice(ctx context.Context, in *InvoiceReq, opts ...grpc.CallOption) (*InvoiceRes, error) {
	out := new(InvoiceRes)
	err := c.cc.Invoke(ctx, "/job.JobService/SendWeeklyInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobProposals(ctx context.Context, in *GJPReq, opts ...grpc.CallOption) (*GJPRes, error) {
	out := new(GJPRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetJobProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobOffersForFreelancer(ctx context.Context, in *GetJobOfferForFreelancerReq, opts ...grpc.CallOption) (*GetJobOfferForFreelancerRes, error) {
	out := new(GetJobOfferForFreelancerRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetJobOffersForFreelancer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SearchJobs(ctx context.Context, in *SearchJobsReq, opts ...grpc.CallOption) (*SearchJobsRes, error) {
	out := new(SearchJobsRes)
	err := c.cc.Invoke(ctx, "/job.JobService/SearchJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetAllContractsForClient(ctx context.Context, in *GetAllContractsForClientReq, opts ...grpc.CallOption) (*GetAllContractsForClientRes, error) {
	out := new(GetAllContractsForClientRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetAllContractsForClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetOneContractForClient(ctx context.Context, in *GetOneContractForClientReq, opts ...grpc.CallOption) (*GetOneContractForClientRes, error) {
	out := new(GetOneContractForClientRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetOneContractForClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ExecutePaymentContract(ctx context.Context, in *ExecutePaymentReq, opts ...grpc.CallOption) (*ExecutePaymentRes, error) {
	out := new(ExecutePaymentRes)
	err := c.cc.Invoke(ctx, "/job.JobService/ExecutePaymentContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) CapturePaymentContract(ctx context.Context, in *CapturePaymentReq, opts ...grpc.CallOption) (*CapturePaymentRes, error) {
	out := new(CapturePaymentRes)
	err := c.cc.Invoke(ctx, "/job.JobService/CapturePaymentContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SaveJobs(ctx context.Context, in *SaveJobsReq, opts ...grpc.CallOption) (*SaveJobsRes, error) {
	out := new(SaveJobsRes)
	err := c.cc.Invoke(ctx, "/job.JobService/SaveJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetInvoiceContract(ctx context.Context, in *GetInvoiceContractReq, opts ...grpc.CallOption) (*GetInvoiceContractRes, error) {
	out := new(GetInvoiceContractRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetInvoiceContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) AddAttachmentToContract(ctx context.Context, in *AddAttachmentReq, opts ...grpc.CallOption) (*AddAttachmentRes, error) {
	out := new(AddAttachmentRes)
	err := c.cc.Invoke(ctx, "/job.JobService/AddAttachmentToContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetAttachments(ctx context.Context, in *GetAttachmentReq, opts ...grpc.CallOption) (*GetAttachmentRes, error) {
	out := new(GetAttachmentRes)
	err := c.cc.Invoke(ctx, "/job.JobService/GetAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) CheckInvoiceStatus(ctx context.Context, in *CheckInvoiceStatusReq, opts ...grpc.CallOption) (*CheckInvoiceStatusRes, error) {
	out := new(CheckInvoiceStatusRes)
	err := c.cc.Invoke(ctx, "/job.JobService/CheckInvoiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations must embed UnimplementedJobServiceServer
// for forward compatibility
type JobServiceServer interface {
	PostJob(context.Context, *PostjobReq) (*PostjobRes, error)
	EditJob(context.Context, *EditjobReq) (*EditjobRes, error)
	SendProposal(context.Context, *ProposalReq) (*ProposalRes, error)
	SendOffer(context.Context, *SendOfferReq) (*SendOfferRes, error)
	AcceptOffer(context.Context, *AcceptOfferReq) (*AcceptOfferRes, error)
	ViewContract(context.Context, *ContractReq) (*ViewContractRes, error)
	AddCategory(context.Context, *AddCategoryReq) (*AddCategoryRes, error)
	GetCategory(context.Context, *GetCategoryReq) (*GetCategoryRes, error)
	GetMyJobs(context.Context, *GetMyJobsReq) (*GetMyJobsRes, error)
	GetJob(context.Context, *GetJobReq) (*GetJobRes, error)
	GetJobs(context.Context, *NoParam) (*GetJobsRes, error)
	SendWeeklyInvoice(context.Context, *InvoiceReq) (*InvoiceRes, error)
	GetJobProposals(context.Context, *GJPReq) (*GJPRes, error)
	GetJobOffersForFreelancer(context.Context, *GetJobOfferForFreelancerReq) (*GetJobOfferForFreelancerRes, error)
	SearchJobs(context.Context, *SearchJobsReq) (*SearchJobsRes, error)
	GetAllContractsForClient(context.Context, *GetAllContractsForClientReq) (*GetAllContractsForClientRes, error)
	GetOneContractForClient(context.Context, *GetOneContractForClientReq) (*GetOneContractForClientRes, error)
	ExecutePaymentContract(context.Context, *ExecutePaymentReq) (*ExecutePaymentRes, error)
	CapturePaymentContract(context.Context, *CapturePaymentReq) (*CapturePaymentRes, error)
	SaveJobs(context.Context, *SaveJobsReq) (*SaveJobsRes, error)
	GetInvoiceContract(context.Context, *GetInvoiceContractReq) (*GetInvoiceContractRes, error)
	AddAttachmentToContract(context.Context, *AddAttachmentReq) (*AddAttachmentRes, error)
	GetAttachments(context.Context, *GetAttachmentReq) (*GetAttachmentRes, error)
	CheckInvoiceStatus(context.Context, *CheckInvoiceStatusReq) (*CheckInvoiceStatusRes, error)
	mustEmbedUnimplementedJobServiceServer()
}

// UnimplementedJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (UnimplementedJobServiceServer) PostJob(context.Context, *PostjobReq) (*PostjobRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostJob not implemented")
}
func (UnimplementedJobServiceServer) EditJob(context.Context, *EditjobReq) (*EditjobRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditJob not implemented")
}
func (UnimplementedJobServiceServer) SendProposal(context.Context, *ProposalReq) (*ProposalRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}
func (UnimplementedJobServiceServer) SendOffer(context.Context, *SendOfferReq) (*SendOfferRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOffer not implemented")
}
func (UnimplementedJobServiceServer) AcceptOffer(context.Context, *AcceptOfferReq) (*AcceptOfferRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOffer not implemented")
}
func (UnimplementedJobServiceServer) ViewContract(context.Context, *ContractReq) (*ViewContractRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewContract not implemented")
}
func (UnimplementedJobServiceServer) AddCategory(context.Context, *AddCategoryReq) (*AddCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedJobServiceServer) GetCategory(context.Context, *GetCategoryReq) (*GetCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedJobServiceServer) GetMyJobs(context.Context, *GetMyJobsReq) (*GetMyJobsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyJobs not implemented")
}
func (UnimplementedJobServiceServer) GetJob(context.Context, *GetJobReq) (*GetJobRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedJobServiceServer) GetJobs(context.Context, *NoParam) (*GetJobsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobs not implemented")
}
func (UnimplementedJobServiceServer) SendWeeklyInvoice(context.Context, *InvoiceReq) (*InvoiceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWeeklyInvoice not implemented")
}
func (UnimplementedJobServiceServer) GetJobProposals(context.Context, *GJPReq) (*GJPRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobProposals not implemented")
}
func (UnimplementedJobServiceServer) GetJobOffersForFreelancer(context.Context, *GetJobOfferForFreelancerReq) (*GetJobOfferForFreelancerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobOffersForFreelancer not implemented")
}
func (UnimplementedJobServiceServer) SearchJobs(context.Context, *SearchJobsReq) (*SearchJobsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchJobs not implemented")
}
func (UnimplementedJobServiceServer) GetAllContractsForClient(context.Context, *GetAllContractsForClientReq) (*GetAllContractsForClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContractsForClient not implemented")
}
func (UnimplementedJobServiceServer) GetOneContractForClient(context.Context, *GetOneContractForClientReq) (*GetOneContractForClientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneContractForClient not implemented")
}
func (UnimplementedJobServiceServer) ExecutePaymentContract(context.Context, *ExecutePaymentReq) (*ExecutePaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutePaymentContract not implemented")
}
func (UnimplementedJobServiceServer) CapturePaymentContract(context.Context, *CapturePaymentReq) (*CapturePaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CapturePaymentContract not implemented")
}
func (UnimplementedJobServiceServer) SaveJobs(context.Context, *SaveJobsReq) (*SaveJobsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveJobs not implemented")
}
func (UnimplementedJobServiceServer) GetInvoiceContract(context.Context, *GetInvoiceContractReq) (*GetInvoiceContractRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoiceContract not implemented")
}
func (UnimplementedJobServiceServer) AddAttachmentToContract(context.Context, *AddAttachmentReq) (*AddAttachmentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAttachmentToContract not implemented")
}
func (UnimplementedJobServiceServer) GetAttachments(context.Context, *GetAttachmentReq) (*GetAttachmentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachments not implemented")
}
func (UnimplementedJobServiceServer) CheckInvoiceStatus(context.Context, *CheckInvoiceStatusReq) (*CheckInvoiceStatusRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInvoiceStatus not implemented")
}
func (UnimplementedJobServiceServer) mustEmbedUnimplementedJobServiceServer() {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_PostJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostjobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PostJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/PostJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PostJob(ctx, req.(*PostjobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_EditJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditjobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).EditJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/EditJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).EditJob(ctx, req.(*EditjobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SendProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SendProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SendProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SendProposal(ctx, req.(*ProposalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SendOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOfferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SendOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SendOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SendOffer(ctx, req.(*SendOfferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_AcceptOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptOfferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).AcceptOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/AcceptOffer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).AcceptOffer(ctx, req.(*AcceptOfferReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ViewContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ViewContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/ViewContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ViewContract(ctx, req.(*ContractReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/AddCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).AddCategory(ctx, req.(*AddCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetCategory(ctx, req.(*GetCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetMyJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyJobsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetMyJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetMyJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetMyJobs(ctx, req.(*GetMyJobsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*GetJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobs(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SendWeeklyInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SendWeeklyInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SendWeeklyInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SendWeeklyInvoice(ctx, req.(*InvoiceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GJPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetJobProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobProposals(ctx, req.(*GJPReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobOffersForFreelancer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobOfferForFreelancerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobOffersForFreelancer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetJobOffersForFreelancer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobOffersForFreelancer(ctx, req.(*GetJobOfferForFreelancerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SearchJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchJobsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SearchJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SearchJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SearchJobs(ctx, req.(*SearchJobsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetAllContractsForClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllContractsForClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetAllContractsForClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetAllContractsForClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetAllContractsForClient(ctx, req.(*GetAllContractsForClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetOneContractForClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneContractForClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetOneContractForClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetOneContractForClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetOneContractForClient(ctx, req.(*GetOneContractForClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ExecutePaymentContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ExecutePaymentContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/ExecutePaymentContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ExecutePaymentContract(ctx, req.(*ExecutePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_CapturePaymentContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapturePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CapturePaymentContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/CapturePaymentContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CapturePaymentContract(ctx, req.(*CapturePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SaveJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveJobsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SaveJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SaveJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SaveJobs(ctx, req.(*SaveJobsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetInvoiceContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoiceContractReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetInvoiceContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetInvoiceContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetInvoiceContract(ctx, req.(*GetInvoiceContractReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_AddAttachmentToContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAttachmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).AddAttachmentToContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/AddAttachmentToContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).AddAttachmentToContract(ctx, req.(*AddAttachmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetAttachments(ctx, req.(*GetAttachmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_CheckInvoiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckInvoiceStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CheckInvoiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/CheckInvoiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CheckInvoiceStatus(ctx, req.(*CheckInvoiceStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostJob",
			Handler:    _JobService_PostJob_Handler,
		},
		{
			MethodName: "EditJob",
			Handler:    _JobService_EditJob_Handler,
		},
		{
			MethodName: "SendProposal",
			Handler:    _JobService_SendProposal_Handler,
		},
		{
			MethodName: "SendOffer",
			Handler:    _JobService_SendOffer_Handler,
		},
		{
			MethodName: "AcceptOffer",
			Handler:    _JobService_AcceptOffer_Handler,
		},
		{
			MethodName: "ViewContract",
			Handler:    _JobService_ViewContract_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _JobService_AddCategory_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _JobService_GetCategory_Handler,
		},
		{
			MethodName: "GetMyJobs",
			Handler:    _JobService_GetMyJobs_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "GetJobs",
			Handler:    _JobService_GetJobs_Handler,
		},
		{
			MethodName: "SendWeeklyInvoice",
			Handler:    _JobService_SendWeeklyInvoice_Handler,
		},
		{
			MethodName: "GetJobProposals",
			Handler:    _JobService_GetJobProposals_Handler,
		},
		{
			MethodName: "GetJobOffersForFreelancer",
			Handler:    _JobService_GetJobOffersForFreelancer_Handler,
		},
		{
			MethodName: "SearchJobs",
			Handler:    _JobService_SearchJobs_Handler,
		},
		{
			MethodName: "GetAllContractsForClient",
			Handler:    _JobService_GetAllContractsForClient_Handler,
		},
		{
			MethodName: "GetOneContractForClient",
			Handler:    _JobService_GetOneContractForClient_Handler,
		},
		{
			MethodName: "ExecutePaymentContract",
			Handler:    _JobService_ExecutePaymentContract_Handler,
		},
		{
			MethodName: "CapturePaymentContract",
			Handler:    _JobService_CapturePaymentContract_Handler,
		},
		{
			MethodName: "SaveJobs",
			Handler:    _JobService_SaveJobs_Handler,
		},
		{
			MethodName: "GetInvoiceContract",
			Handler:    _JobService_GetInvoiceContract_Handler,
		},
		{
			MethodName: "AddAttachmentToContract",
			Handler:    _JobService_AddAttachmentToContract_Handler,
		},
		{
			MethodName: "GetAttachments",
			Handler:    _JobService_GetAttachments_Handler,
		},
		{
			MethodName: "CheckInvoiceStatus",
			Handler:    _JobService_CheckInvoiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/Job/job.proto",
}
