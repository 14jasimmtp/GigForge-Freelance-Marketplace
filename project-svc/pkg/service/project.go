package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb/user"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/utils/paypal"
)

type ProjectService struct{
	pb.UnimplementedProjectServiceServer
	repo repository.Repo
	user user.UserServiceClient
}

func NewService(rep repository.Repo,user user.UserServiceClient) ProjectService{
	return ProjectService{repo: rep,user: user}
}

func (s *ProjectService) AddProject(ctx context.Context,req *pb.AddSingleProjectReq) (*pb.AddSingleProjectRes,error){
	status,err:=s.repo.AddSingleProject(req)
	if err != nil {
		return &pb.AddSingleProjectRes{
			Status: int64(status),
			Error: err.Error(),
		},nil
	}
	return &pb.AddSingleProjectRes{
		Status: int64(status),
		Response: "added project successfully",
	},nil
}

func (s *ProjectService) EditProject(ctx context.Context,req *pb.EditSingleProjectReq) (*pb.EditSingleProjectRes,error){
	status,err:=s.repo.EditSingleProject(req)
	if err != nil {
		return &pb.EditSingleProjectRes{
			Status: int64(status),
			Error: err.Error(),
		},nil
	}
	return &pb.EditSingleProjectRes{
		Status: int64(status),
		Response: "updated project successfully",
	},nil
}

func (s *ProjectService) RemoveProject(ctx context.Context, req *pb.RemProjectReq) (*pb.RemProjectRes,error){
	err:=s.repo.DeleteProject(req)
	if err != nil {
		return &pb.RemProjectRes{Status: 400,Error: err.Error()},nil
	}
	return &pb.RemProjectRes{Status: 200,Response: "project deleted successfully"},nil
}

func (s *ProjectService)  ListProjects(ctx context.Context, req *pb.NoParam) (*pb.ListProjectsRes,error){
	projects,err:=s.repo.ListProjects()
	if err != nil {
		return &pb.ListProjectsRes{Status: 400,Error: err.Error()},nil

	}
	return &pb.ListProjectsRes{Project: projects,Status: 200,Response: "projects fetched successfully"},nil
}

func (s *ProjectService) ListOneProject(ctx context.Context, req *pb.ListOneProjectReq) (*pb.ListOneProjectRes,error){
	project,err:=s.repo.ListOneProject(req.ProjectId)
	if err != nil {
		return &pb.ListOneProjectRes{Status: 400,Error: err.Error()},nil

	}
	return &pb.ListOneProjectRes{Project: project,Status: 200,Response: "project fetched successfully"},nil
}

func (s *ProjectService) ListMyProjects(ctx context.Context, req *pb.ListMyProjectReq) (*pb.ListMyProjectRes,error){
	projects,err:=s.repo.ListMyProject(req.UserId)
	if err != nil {
		return &pb.ListMyProjectRes{Status: 400,Error: err.Error()},nil

	}
	return &pb.ListMyProjectRes{Project: projects,Status: 200,Response: "project fetched successfully"},nil
}

func (s *ProjectService) ExecutePaymentProject(ctx context.Context, req *pb.ExecutePaymentReq) (*pb.ExecutePaymentRes, error) {
	fmt.Println("executing payment")
	order, err := s.repo.GetProjectOrder(req.OrderID)
	if err != nil {
		return &pb.ExecutePaymentRes{Status: http.StatusExpectationFailed, Error: err.Error()}, nil
	}
	if order.Payment_status == "paid"{
		return &pb.ExecutePaymentRes{Status: http.StatusExpectationFailed,Error: "already paid for the project"},nil
	}
	freelancerEmail, err := s.user.GetFreelancerPaypalEmails(context.Background(), &user.Preq{UserID: int32(order.FreelancerID)})
	if err != nil {
		return &pb.ExecutePaymentRes{Status: http.StatusBadRequest, Error: freelancerEmail.Error}, nil
	}
	orders, err := paypal.CreatePayment(order, freelancerEmail.Email)
	if err != nil {
		return &pb.ExecutePaymentRes{Status: http.StatusFailedDependency, Error: err.Error()}, nil
	}
	return &pb.ExecutePaymentRes{Status: http.StatusOK, PaymentID: orders.OrderID, MerchantID: orders.MerchantID}, nil
}

func (s *ProjectService) CapturePaymentContract(ctx context.Context, req *pb.CapturePaymentReq) (*pb.CapturePaymentRes, error) {
	fmt.Println("capturing payment...")
	ClientName, err := paypal.CapturePayment(req.PaymentID)
	if err != nil {
		return &pb.CapturePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	_, err = s.repo.UpdateOrderPaymentStatus(req.OrderID)
	if err != nil {
		return &pb.CapturePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	return &pb.CapturePaymentRes{Status: http.StatusOK, UserName: ClientName}, nil
}

 func (s *ProjectService) OrderProject(ctx context.Context,req *pb.BuyProjectReq) (*pb.BuyProjectRes,error){
	project,pro,err:=s.repo.CheckProjectActiveAndExist(req.ProjectId)
	if err != nil {
		return &pb.BuyProjectRes{Status: http.StatusForbidden,Error: "project is not currently active."},nil
	}
	err=s.repo.OrderProject(project,pro,req.UserId)
	if err != nil {
		return &pb.BuyProjectRes{Status: http.StatusInternalServerError,Error: err.Error()},nil
	}
	return &pb.BuyProjectRes{Status: http.StatusOK,Response: "project order successful. Will attach to you within delivery time."},nil
 }

