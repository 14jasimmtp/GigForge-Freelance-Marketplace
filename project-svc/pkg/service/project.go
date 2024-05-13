package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/project-svc/pkg/repository"
)

type ProjectService struct{
	pb.UnimplementedProjectServiceServer
	repo repository.Repo
}

func NewService(rep repository.Repo) ProjectService{
	return ProjectService{repo: rep}
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

// func (s *ProjectService) PaymentForProject(ctx context.Context, req *pb.ProjectPaymentReq) (*pb.ProjectPaymentRes,error){

// }

//  func (s *ProjectService) BuyProject(ctx context.Context,req *pb.BuyProjectReq) (*pb.BuyProjectRes,error){
// 	err:=s.repo.OrderProject(req.ProjectId,req.UserId)
// 	if err != nil {
// 		return &pb.BuyProjectRes{Status: 400,Error: err.Error()},nil
// 	}

//  }

