package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/repository"
)

type Service struct{
	job.UnimplementedJobServiceServer
	repo repository.Repo
}

func NewJobService(repo repository.Repo) Service{
	return Service{repo: repo}
}

func (s *Service) PostJob(ctx context.Context, req *job.PostjobReq)(*job.PostjobRes,error){
	err:=s.repo.PostJob(req)
	if err != nil {
		return &job.PostjobRes{
			Status: 500,
			Error: err.Error(),
			Response: "error while adding job to database",
		},nil
	}
	return &job.PostjobRes{
		Status: 200,
		Response: "job posted successfully",
	},nil
}

func (s *Service) SendProposal(ctx context.Context, req *job.ProposalReq) (*job.ProposalRes,error){
	err:=s.repo.FindJob(req.JobId)
	if err != nil {
		return &job.ProposalRes{
			Status: 500,
			Error: err.Error(),
			Response: "No job found with this id",
		},nil
	}
	err =s.repo.Proposal(req)
	if err != nil {
		return &job.ProposalRes{
			Status: 500,
			Error: err.Error(),
			Response: "something went wrong",
		},nil
	}
	return &job.ProposalRes{
		Status: 200,
		Response: "proposal sent successfully",
	},nil
}

func (s *Service) SendOffer(ctx context.Context, req *job.SendOfferReq)(*job.SendOfferRes,error){
	err := 
}

func (s *Service) AcceptOffer()



// func (s *Service) AcceptProposal(ctx context.Context, req *job.AcceptOfferReq) (*job.AcceptOfferRes, error) {
// 	err:=
// }