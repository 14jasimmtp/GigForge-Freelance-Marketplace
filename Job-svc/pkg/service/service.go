package service

import (
	"context"
	"net/http"
	"strconv"
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

func (s *Service) GetJobProposals(ctx context.Context,req *job.GJPReq) (*job.GJPRes,error){
	err:=s.repo.FindJobExistOfClient(req.JobId,req.UserId)
	if err != nil {
		return &job.GJPRes{Status: http.StatusBadRequest,Error: err.Error()},nil
	}
	proposals, err:=s.repo.GetJobProposals(req.JobId)
	if err != nil {
		return &job.GJPRes{Status: http.StatusBadRequest, Error: err.Error()},nil
	}
	return &job.GJPRes{Status: http.StatusOK,Prop: proposals,Response: "fetched proposals successfully"},nil
}

func (s *Service) GetCategories(ctx context.Context,req *job.GetCategoryReq) (*job.GetCategoryRes,error){
	categories,err:=s.repo.GetCategory(req.Query)
	if err != nil {
		return &job.GetCategoryRes{Status: http.StatusNoContent,Error: err.Error()},nil
	}
	return &job.GetCategoryRes{Status: http.StatusOK,Categories: categories},nil
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
	_,err:=time.Parse("02-01-2006",req.StartingTime)
	if err != nil {
		return &job.SendOfferRes{Status: http.StatusBadRequest,Error: err.Error()},nil
	}
	// s.repo.CheckJobExist(req.JobId)
	// s.repo.CheckFreelancerExist(req.FreelancerId)
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
	if err != nil {
		return &job.AcceptOfferRes{Status: http.StatusBadRequest,Error: err.Error()},nil
	}
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
	contract, err := s.repo.CheckContractActive(req.ContractID)
	if err != nil {
		return &job.InvoiceRes{Status: 400, Error: err.Error()}, nil
	}
	if strconv.Itoa(contract.Freelancer_id) != req.SuserId {
		return &job.InvoiceRes{Status: 400, Error: "Not your contract. Check Contract ID is correct"}, nil
	}
	LastInvoice, err := s.repo.GetLastInvoiceDetails(req.ContractID)
	if err != nil {
		return &job.InvoiceRes{Status: 500, Error: err.Error()}, nil
	}
	StartDate, err := time.Parse("02-01-2006", req.StartDate)
	if err != nil {
		return &job.InvoiceRes{Status: 400, Error: err.Error()}, nil
	}
	EndDate, err := time.Parse("02-01-2006", req.EndDate)
	if err != nil {
		return &job.InvoiceRes{Status: 400, Error: err.Error()}, nil
	}
	if LastInvoice.End_date.After(StartDate) {
		return &job.InvoiceRes{Status: http.StatusConflict, Error: "start date already covered in last invoice.Give the correct start date"}, nil
	}

	if EndDate.After(StartDate) {
		return &job.InvoiceRes{Status: http.StatusBadRequest, Error: "End date is before start date. Enter date correctly"}, nil
	}

	err = s.repo.SendHourlyInvoice(int(contract.ID), contract.Type, contract.Budget, req.TotalHourWorked, EndDate, StartDate)
	if err != nil {
		return &job.InvoiceRes{Status: 500, Error: err.Error()}, nil
	}
	return &job.InvoiceRes{Status: 200, Response: "invoice sent successfully"}, nil
}

func (s *Service) SearchJobs(ctx context.Context, req *job.SearchJobsReq) (*job.SearchJobsRes, error) {
	FixedRate, HourlyRate := []string{}, []string{}
	if req.FixedRate != "" {
		FixedRate = strings.Split(req.FixedRate, "-")
	} else {

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

func (s *Service) GetJobOffersForFreelancer(ctx context.Context,req *job.GetJobOfferForFreelancerReq) (*job.GetJobOfferForFreelancerRes,error){
	offers,err:=s.repo.GetOffers(req.UserId,req.Status) 
	if err != nil {
		return &job.GetJobOfferForFreelancerRes{Status: http.StatusBadRequest,Error: err.Error()},nil
	}
	return &job.GetJobOfferForFreelancerRes{Status: http.StatusOK,Offers: offers},nil
}

//  func (s *Service) ExecutePayment(ctx context.Context,req *job.ExecutePaymentReq)(*job.ExecutePaymentRes,error){
// 	invoice,err:=s.repo.GetInvoiceDetails(req.InvoiceId)
// 	if err != nil {
// 		return &job.ExecutePaymentRes{},nil
// 	}
// 	accessToken,err:=paypal.GenerateAccessToken()
// 	auth_assertion_header:=paypal.GetAuthAssertionValue(req.User_id)
// 	OrderID,err:=paypal.CreateOrder(accessToken,invoice.ID,invoice.Budget,"USD",auth_assertion_header)
// 	if err != nil {
// 		return &job.ExecutePaymentRes{},nil
// 	}
// 	return &job.ExecutePaymentRes{}

// }

// }
