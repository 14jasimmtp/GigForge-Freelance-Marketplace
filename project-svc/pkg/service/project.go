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

// func (s *ProjectService) AddTieredProject(ctx context.Context, req *pb.AddProjectReq) (*pb.AddProjectRes){

// }

// func (s *ProjectService) GetCheckout(ctx context.Context){

// }

// func (s *ProjectService) BuyProject(ctx context.Context)

// func (s *ProjectService) ViewProject(ctx context.Context)