package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/user"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/utils/paypal"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/utils/s3"
)

type Service struct {
	job.UnimplementedJobServiceServer
	repo repository.Repo
	user user.JobserviceClient
}

func NewJobService(repo repository.Repo, user user.JobserviceClient) *Service {
	return &Service{repo: repo, user: user}
}

func (s *Service) PostJob(ctx context.Context, req *job.PostjobReq) (*job.PostjobRes, error) {
	res,err:=s.user.CheckPaypalEmailAdded(context.Background(),&user.CReq{UserId: req.UserId})
	if err != nil {
		return &job.PostjobRes{Status: int64(res.Status),Error: "error fetching user paypal"},nil
	}
	if !res.Exist{
		return &job.PostjobRes{Status: int64(res.Status),Error: "add payment email before adding a job post"},nil
	}

	err = s.repo.PostJob(req)
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

func (s *Service) EditJob(ctx context.Context, req *job.EditjobReq) (*job.EditjobRes, error){
	err:=s.repo.CheckJobExist(req.JobID,req.UserId)
	if err != nil {
		return &job.EditjobRes{	Error: err.Error(),Status: http.StatusNotFound},nil
	}

	err = s.repo.EditJobPost(req)
	if err != nil {
		return &job.EditjobRes{	Error: err.Error(),Status: http.StatusFailedDependency},nil
	}
	return &job.EditjobRes{Response: "Job updated successfully",Status: http.StatusOK},nil
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

func (s *Service) GetJobProposals(ctx context.Context, req *job.GJPReq) (*job.GJPRes, error) {
	err := s.repo.FindJobExistOfClient(req.JobId, req.UserId)
	if err != nil {
		return &job.GJPRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	proposals, err := s.repo.GetJobProposals(req.JobId)
	if err != nil {
		return &job.GJPRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	return &job.GJPRes{Status: http.StatusOK, Prop: proposals, Response: "fetched proposals successfully"}, nil
}

func (s *Service) GetCategories(ctx context.Context, req *job.GetCategoryReq) (*job.GetCategoryRes, error) {
	categories, err := s.repo.GetCategory(req.Query)
	if err != nil {
		return &job.GetCategoryRes{Status: http.StatusNoContent, Error: err.Error()}, nil
	}
	return &job.GetCategoryRes{Status: http.StatusOK, Categories: categories}, nil
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
	_, err := time.Parse("02-01-2006", req.StartingTime)
	if err != nil {
		return &job.SendOfferRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	err = s.repo.CheckJobExist(strconv.Itoa(int(req.JobId)),int64(req.ClientId))
	if err != nil {
		return &job.SendOfferRes{Status: http.StatusBadRequest,Error: err.Error()},nil
	}
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
	uid,err:=strconv.Atoi(req.UserId)
	res,err := s.user.CheckPaypalEmailAdded(context.Background(),&user.CReq{UserId: int64(uid)})
	if err != nil {
		return &job.AcceptOfferRes{Status: http.StatusInternalServerError,Error: err.Error()},nil
	}
	if !res.Exist{
		return &job.AcceptOfferRes{Status: http.StatusBadRequest,Error: res.Error},nil
	}

	err = s.repo.AcceptOffer(req.OfferID)
	if err != nil {
		return &job.AcceptOfferRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	contractID, contractType, budget, err := s.repo.CreateContract(req.OfferID)
	if err != nil {
		return &job.AcceptOfferRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
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

func (s *Service) GetJobOffersForFreelancer(ctx context.Context, req *job.GetJobOfferForFreelancerReq) (*job.GetJobOfferForFreelancerRes, error) {
	offers, err := s.repo.GetOffers(req.UserId, req.Status)
	if err != nil {
		return &job.GetJobOfferForFreelancerRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	return &job.GetJobOfferForFreelancerRes{Status: http.StatusOK, Offers: offers}, nil
}

func (s *Service) ExecutePaymentContract(ctx context.Context, req *job.ExecutePaymentReq) (*job.ExecutePaymentRes, error) {
	fmt.Println("executing payment")
	invoice, err := s.repo.GetInvoiceWithID(req.InvoiceId)
	if err != nil {
		return &job.ExecutePaymentRes{Status: http.StatusExpectationFailed, Error: err.Error()}, nil
	}
	contract, err := s.repo.CheckContractActive(int32(invoice.ContractID))
	if err != nil {
		return &job.ExecutePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	freelancerEmail, err := s.user.GetFreelancerPaypalEmail(context.Background(), &user.Preq{UserID: int32(contract.Freelancer_id)})
	if err != nil {
		return &job.ExecutePaymentRes{Status: http.StatusBadRequest, Error: freelancerEmail.Error}, nil
	}
	order, err := paypal.CreatePayment(invoice, freelancerEmail.Email)
	if err != nil {
		return &job.ExecutePaymentRes{Status: http.StatusFailedDependency, Error: err.Error()}, nil
	}
	return &job.ExecutePaymentRes{Status: http.StatusOK, PaymentID: order.OrderID, MerchantID: order.MerchantID}, nil
}

func (s *Service) CapturePaymentContract(ctx context.Context, req *job.CapturePaymentReq) (*job.CapturePaymentRes, error) {
	fmt.Println("capturing payment...")
	ClientName, err := paypal.CapturePayment(req.PaymentID)
	if err != nil {
		return &job.CapturePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	invoice, err := s.repo.UpdateInvoicePaymentStatus(req.InvoiceID)
	if err != nil {
		return &job.CapturePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	err = s.repo.UpdateContractDetails(invoice.ContractID, invoice.Freelancer_fee, invoice.MarketPlace_fee)
	if err != nil {
		return &job.CapturePaymentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	return &job.CapturePaymentRes{Status: http.StatusOK, UserName: ClientName}, nil
}

func (s *Service) GetAllContractsForClient(ctx context.Context, req *job.GetAllContractsForClientReq) (*job.GetAllContractsForClientRes, error) {
	contracts, err := s.repo.GetAllContracts(req.UserId)
	if err != nil {
		return &job.GetAllContractsForClientRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	var con []*job.Contracts
	for _, c := range contracts {
		Jobs, err := s.repo.GetJob(fmt.Sprintf("%d", c.Job_id))
		if err != nil {
			return &job.GetAllContractsForClientRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
		}
		con = append(con, &job.Contracts{
			ContractId:     int32(c.ID),
			FreelancerId:   int32(c.Freelancer_id),
			ClientId:       int32(c.Client_id),
			PaymentType:    c.Type,
			TotalAmount:    c.Budget,
			PaidAmount:     float32(c.Paid_amount),
			PendingAmount:  float32(c.Pending_amount),
			ContractStatus: c.Status,
			StartDate:      c.Start_date.Format("02-01-2006"),
			JobTitle:       Jobs.Title,
			JobDescription: Jobs.Description,
		})
	}

	return &job.GetAllContractsForClientRes{Contracts: con, Status: http.StatusOK}, nil
}

func (s *Service) GetOneContractForClient(ctx context.Context, req *job.GetOneContractForClientReq) (*job.GetOneContractForClientRes, error) {
	contract, err := s.repo.GetOneContract(req.ContractID, req.UserId)
	if err != nil {
		return &job.GetOneContractForClientRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	Jobs, err := s.repo.GetJob(fmt.Sprintf("%d", contract.Job_id))
	if err != nil {
		return &job.GetOneContractForClientRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	con := &job.Contracts{
		ContractId:     int32(contract.ID),
		FreelancerId:   int32(contract.Freelancer_id),
		ClientId:       int32(contract.Client_id),
		PaymentType:    contract.Type,
		TotalAmount:    contract.Budget,
		PaidAmount:     float32(contract.Paid_amount),
		PendingAmount:  float32(contract.Pending_amount),
		ContractStatus: contract.Status,
		StartDate:      contract.Start_date.Format("02-01-2006"),
		JobTitle:       Jobs.Title,
		JobDescription: Jobs.Description,
	}
	return &job.GetOneContractForClientRes{Status: http.StatusOK, Contract: con}, nil

}

func (s *Service) GetInvoiceContract(ctx context.Context, req *job.GetInvoiceContractReq) (*job.GetInvoiceContractRes, error) {
	invoices, err := s.repo.GetInvoices(req.UserID, req.ContractID)
	if err != nil {
		return &job.GetInvoiceContractRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}
	var inv []*job.Invoices
	for _, i := range invoices {
		in := &job.Invoices{
			InvoiceID:      int32(i.ID),
			ContractID:     int32(i.ContractID),
			StartDate:      i.Start_date.Format("02-01-2006"),
			EndDate:        i.End_date.Format("02-01-2006"),
			PaymentStatus:  i.Status,
			FreelancerFee:  float32(i.Freelancer_fee),
			MarketPlaceFee: float32(i.MarketPlace_fee),
		}
		inv = append(inv, in)
	}
	return &job.GetInvoiceContractRes{Invoices: inv, Status: http.StatusOK}, nil
}

func (s *Service) AddAttachmentToContract(ctx context.Context, req *job.AddAttachmentReq) (*job.AddAttachmentRes, error) {
	cid, _ := strconv.Atoi(req.ContractID)
	_, err := s.repo.CheckContractActive(int32(cid))
	if err != nil {
		return &job.AddAttachmentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil
	}

	sess := s3.CreateSession()

	attachmentUrl, err := s3.UploadAttachmentToS3(req.Attachment, req.ContractID, req.Filename, sess)
	if err != nil {
		return &job.AddAttachmentRes{Status: http.StatusBadRequest, Error: `something went wrong while uploading attachment`}, nil
	}
	err = s.repo.StoreAttachmentUrl(attachmentUrl, req)
	if err != nil {
		return &job.AddAttachmentRes{Status: http.StatusBadRequest, Error: `something went wrong while uploading attachment`}, nil

	}
	return &job.AddAttachmentRes{Status: http.StatusOK, Response: "attachment added successfully"}, nil
}


func (s *Service) GetAttachments(ctx context.Context, req *job.GetAttachmentReq) ( *job.GetAttachmentRes,error){
	attachments,err:=s.repo.GetAttachments(req.ContractID)
	if err != nil {
		return &job.GetAttachmentRes{Status: http.StatusBadRequest, Error: err.Error()}, nil

	} 
	return &job.GetAttachmentRes{Status: http.StatusOK,Attachment: attachments},nil
}