package service

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/repository"
)

type Service struct {
	job.UnimplementedJobServiceServer
	repo repository.Repo
}

func NewJobService(repo repository.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) PostJob(ctx context.Context, req *job.PostjobReq) (*job.PostjobRes, error) {
	err := s.repo.PostJob(req)
	if err != nil {
		return &job.PostjobRes{
			Status:   500,
			Error:    err.Error(),
			Response: "error while adding job to database",
		}, nil
	}
	return &job.PostjobRes{
		Status:   200,
		Response: "job posted successfully",
	}, nil
}

func (s *Service) GetMyJobs(ctx context.Context, req *job.GetMyJobsReq) (*job.GetMyJobsRes, error) {
	res, err := s.repo.GetMyJobs(req.UserId)
	if err != nil {
		return &job.GetMyJobsRes{Status: 400, Error: err.Error()}, nil
	}
	return &job.GetMyJobsRes{Status: 200, Jobs: res}, nil
}

func (s *Service) GetJobs(ctx context.Context, req *job.NoParam) (*job.GetJobsRes, error) {
	res, err := s.repo.GetJobs()
	if err != nil {
		return &job.GetJobsRes{Status: 400, Error: err.Error()}, nil
	}
	return &job.GetJobsRes{Status: 200, Job: res}, nil
}

func (s *Service) GetJob(ctx context.Context, req *job.GetJobReq) (*job.GetJobRes, error) {
	res, err := s.repo.GetJob(req.JobId)
	if err != nil {
		return &job.GetJobRes{Status: 400, Error: err.Error()}, nil
	}
	return &job.GetJobRes{Status: 200, Job: res}, nil
}

func (s *Service) SendProposal(ctx context.Context, req *job.ProposalReq) (*job.ProposalRes, error) {
	err := s.repo.FindJob(req.JobId)
	if err != nil {
		return &job.ProposalRes{
			Status:   500,
			Error:    err.Error(),
			Response: "No job found with this id",
		}, nil
	}
	err = s.repo.Proposal(req)
	if err != nil {
		return &job.ProposalRes{
			Status:   500,
			Error:    err.Error(),
			Response: "something went wrong",
		}, nil
	}
	return &job.ProposalRes{
		Status:   200,
		Response: "proposal sent successfully",
	}, nil
}

func (s *Service) SendOffer(ctx context.Context, req *job.SendOfferReq) (*job.SendOfferRes, error) {
	res, err := s.repo.SendOffer(req)
	if err != nil {
		return &job.SendOfferRes{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}
	return res, nil
}

func (s *Service) AcceptOffer(ctx context.Context, req *job.AcceptOfferReq) (*job.AcceptOfferRes, error) {
	err := s.repo.AcceptOffer(req.OfferID)
	if err != nil {
		return &job.AcceptOfferRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	contractID, contractType, budget, err := s.repo.CreateContract(req.OfferID)
	if contractType == "fixed" {
		err := s.repo.SendFixedInvoice(contractID, contractType, budget)
		if err != nil {
			return &job.AcceptOfferRes{
				Status: 400,
				Error:  err.Error(),
			}, nil
		}
		return &job.AcceptOfferRes{Status: 200, Response: "Accepted Offer. Contract created and payment invoice sent to client"}, nil
	}
	return &job.AcceptOfferRes{Status: 200, Response: "Accepted Offer. Send Invoices on Weekends to claim the payment"}, nil
}

func (s *Service) SendWeeklyInvoice(ctx context.Context, req *job.InvoiceReq) (*job.InvoiceRes, error) {
	res, err := s.repo.GetContractDetails(req.ContractID)
	if err != nil {
		return &job.InvoiceRes{Status: 500, Error: err.Error()}, nil
	}
	if res.UpdatedAt.Add(24 * 7 * time.Hour).After(time.Now()) {
		err := s.repo.SendHourlyInvoice(int(res.ID), res.Type, res.Budget, req.TotalHourWorked)
		if err != nil {
			return &job.InvoiceRes{Status: 500, Error: err.Error()}, nil
		}
		return &job.InvoiceRes{Status: 200, Response: "invoice sent successfully"}, nil
	}
	return &job.InvoiceRes{Status: 500, Error: "week is not completed to send invoice"}, nil
}

func (s *Service) SearchJobs(ctx context.Context, req *job.SearchJobsReq) (*job.SearchJobsRes, error) {
	FixedRate, HourlyRate := []string{}, []string{}
	if req.FixedRate != "" {
		FixedRate = strings.Split(req.FixedRate, "-")
	}else{
		
	}
	if req.HourlyRate != "" {
		HourlyRate = strings.Split(req.HourlyRate, "-")
	}
	if req.Paytype == "0" {
		req.Paytype = "hourly"
	} else if req.Paytype == "1" {
		req.Paytype = "fixed"
	}
	res, status, err := s.repo.SearchJobs(req.Category, req.Paytype, req.Query, FixedRate, HourlyRate)
	if err != nil {
		return &job.SearchJobsRes{Error: err.Error(), Status: status, Response: "error while fetching jobs"}, nil
	}
	return &job.SearchJobsRes{Status: status, Jobs: res, Response: "fetched jobs successfully"}, nil
}



// func (s *Service) ExecutePayment(ctx context.Context){
// 	helper.GetPaymentDetails()
// 	helper.GetContractDetails()
// 	helper.CreatePaymentOrder()

// }
