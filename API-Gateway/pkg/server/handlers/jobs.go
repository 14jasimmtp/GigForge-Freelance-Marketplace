package handler

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/Job"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
)

type JobsHandler struct {
	job Job.JobServiceClient
}

func NewJobsHandler(job Job.JobServiceClient) *JobsHandler {
	return &JobsHandler{job: job}
}

// PostJob godoc
// @Summary Post a job
// @Description Create a new job listing
// @security ClientAccessToken
// @Tags jobs
// @Accept json
// @Produce json
// @Param job body req.PostJob true "Job details"
// @Success 200 {object} Job.PostjobRes "Successfully posted job"
// @Failure 400 {object} res.CommonRes "Error validating request body"
// @Failure 403 {object} map[string]string "Forbidden"
// @Router /client/job [post]
func (h *JobsHandler) PostJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	res, err := h.job.PostJob(context.Background(), &Job.PostjobReq{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Skills:      req.Skills,
		TimePeriod:  req.TimePeriod,
		Type:        req.Type,
		Budget:      float32(req.Budget),
		UserId:      int64(user_id),
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)

}

// SendProposal godoc
// @Summary Send a proposal
// @Description Send a proposal for a job
// @security FreelancerAccessToken
// @Tags jobs
// @Accept json
// @Produce json
// @Param jobID query string true "Job ID"
// @Param proposal body req.Proposal true "Proposal details"
// @Success 200 {object} Job.ProposalRes "Successfully sent proposal"
// @Failure 400 {object} res.CommonRes "Error validating request body"
// @Failure 403 {object} map[string]string "Forbidden"
// @Router /freelancer/job/send-proposal [post]
func (h *JobsHandler) SendProposal(c *fiber.Ctx) error {
	var req req.Proposal

	job_id := c.Query("jobID")
	user_id := c.Locals("User_id").(int)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	res, err := h.job.SendProposal(context.Background(), &Job.ProposalReq{
		Budget:      req.Budget,
		CoverLetter: req.Coverletter,
		UserId:      strconv.Itoa(user_id),
		JobId:       job_id,
		// Attachments: attachments,
	})

	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// GetProposalsOfJob godoc
// @Summary Get proposals of a job
// @Description Get all proposals for a specific job
// @security ClientAccessToken
// @Tags jobs
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} Job.GJPRes "Successfully retrieved proposals"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /client/job/proposals/{id} [get]
func (h *JobsHandler) GetProposalsOfJob(c *fiber.Ctx) error {
	client_id := strconv.Itoa(int(c.Locals("User_id").(int)))
	job_id := c.Params("id")

	res, err := h.job.GetJobProposals(context.Background(), &Job.GJPReq{UserId: client_id, JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetMyJobs godoc
// @Summary Get my jobs
// @Description Get jobs posted by the authenticated user
// @security ClientAccessToken
// @Tags jobs
// @Produce json
// @Success 200 {object} Job.GetMyJobsRes "Successfully retrieved jobs"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /client/job [get]
func (h *JobsHandler) GetMyJobs(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	id := strconv.Itoa(int(user_id))
	res, err := h.job.GetMyJobs(context.Background(), &Job.GetMyJobsReq{UserId: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetCategories godoc
// @Summary Get categories
// @Description Get job categories based on a query
// @Tags jobs
// @Produce json
// @Param q query string true "Query string"
// @Success 200 {object} Job.GetCategoryRes "Successfully retrieved categories"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /job/categories [get]
func (h *JobsHandler) GetCategories(c *fiber.Ctx) error {
	query := c.Query("q")
	category, err := h.job.GetCategory(context.Background(), &Job.GetCategoryReq{Query: query})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(int(category.Status)).JSON(category)
}

// GetJobProposals godoc
// @Summary Get job proposals
// @Description Get all proposals for a specific job
// @Tags jobs
// @Produce json
// @Param job_id path string true "Job ID"
// @Success 200 {object} Job.GJPRes "Successfully retrieved proposals"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /job/{job_id}/proposals [get]
func (h *JobsHandler) GetJobProposals(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	user_id := c.Locals("User_id").(int)
	res, err := h.job.GetJobProposals(context.Background(), &Job.GJPReq{JobId: jobID, UserId: strconv.Itoa(user_id)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// EditJob godoc
// @Summary Edit a job
// @Description Edit an existing job listing
// @security ClientAccessToken
// @Tags jobs
// @Accept json
// @Produce json
// @Param job body req.PostJob true "Job details"
// @Success 200 {object} Job.PostjobRes "Successfully edited job"
// @Failure 400 {object} res.CommonRes "Error validating request body"
// @Failure 403 {object} map[string]string "Forbidden"
// @Router /client/job [put]
func (h *JobsHandler) EditJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	res, err := h.job.PostJob(context.Background(), &Job.PostjobReq{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Skills:      req.Skills,
		TimePeriod:  req.TimePeriod,
		Type:        req.Type,
		Budget:      float32(req.Budget),
		UserId:      int64(user_id),
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// SendOffer godoc
// @Summary Send a job offer
// @Description Send a job offer to a freelancer
// @security ClientAccessToken
// @Tags jobs
// @Accept json
// @Produce json
// @Param offer body req.SendOffer true "Offer details"
// @Success 200 {object} Job.SendOfferRes "Successfully sent offer"
// @Failure 400 {object} res.CommonRes "Error validating request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /client/job/send-offer [post]
func (h *JobsHandler) SendOffer(c *fiber.Ctx) error {
	var req req.SendOffer
	user_id := c.Locals("User_id").(int)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}
	startDate, err := time.Parse("2-1-2006", req.Starting_time)
	if err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": Error})
	}
	res, err := h.job.SendOffer(context.Background(), &Job.SendOfferReq{
		Budget:       req.Budget,
		OfferLetter:  req.Offer_letter,
		StartingTime: startDate.Format("2-1-2006"),
		JobId:        int32(req.Job_id),
		FreelancerId: int32(req.Freelancer_id),
		ClientId:     int32(user_id),
	})
	if err != nil {
		println(err)
		return c.Status(500).JSON(fiber.Map{"error": "error in rpc connection"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// AcceptOffer godoc
// @Summary Accept a job offer
// @Description Accept a job offer from a client
// @security FreelancerAccessToken
// @Tags jobs
// @Produce json
// @Param offer_id path string true "Offer ID"
// @Success 200 {object} Job.AcceptOfferRes "Successfully accepted offer"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /freelancer/job/accept-offer/{offer_id} [post]
func (h *JobsHandler) AcceptOffer(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	of_id := c.Params("offer_id")

	res, err := h.job.AcceptOffer(context.Background(), &Job.AcceptOfferReq{UserId: strconv.Itoa(user_id), OfferID: of_id})
	if err != nil {
		print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetJobs godoc
// @Summary Get all jobs
// @Description Get a list of all jobs
// @Tags jobs
// @Produce json
// @Success 200 {object} Job.GetJobsRes "Successfully retrieved jobs"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /jobs [get]
func (h *JobsHandler) GetJobs(c *fiber.Ctx) error {
	res, err := h.job.GetJobs(context.Background(), &Job.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetJob godoc
// @Summary Get a job
// @Description Get details of a specific job
// @Tags jobs
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} Job.GetJobRes "Successfully retrieved job"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /job/{id} [get]
func (h *JobsHandler) GetJob(c *fiber.Ctx) error {
	job_id := c.Params("id")
	println(job_id)
	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetAllJobOffersForFreelancer godoc
// @Summary Get job offers for a freelancer
// @Description Get all job offers for a specific freelancer
// @security FreelancerAccessToken
// @Tags jobs
// @Produce json
// @Param status query string false "Offer status"
// @Success 200 {object} Job.GetJobOfferForFreelancerRes "Successfully retrieved job offers"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /freelancer/job/offers [get]
func (h *JobsHandler) GetAllJobOffersForFreelancer(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	status := c.Query("status")
	res, err := h.job.GetJobOffersForFreelancer(context.Background(), &Job.GetJobOfferForFreelancerReq{UserId: strconv.Itoa(user_id), Status: status})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// SendInvoice godoc
// @Summary Send an invoice
// @Description Send a weekly invoice for a contract
// @security FreelancerAccessToken
// @Tags jobs
// @Accept json
// @Produce json
// @Param invoice body req.SendInvoice true "Invoice details"
// @Success 200 {object} Job.InvoiceRes "Successfully sent invoice"
// @Failure 400 {object} res.CommonRes "Error validating request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /freelancer/job/invoice [post]
func (h *JobsHandler) SendInvoice(c *fiber.Ctx) error {
	var req req.SendInvoice
	user_id := c.Locals("User_id").(int)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}
	res, err := h.job.SendWeeklyInvoice(context.Background(), &Job.InvoiceReq{ContractID: int32(req.ContractId), TotalHourWorked: float32(req.TotalHoursWorked), SuserId: strconv.Itoa(user_id)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// ExecutePaymentContract godoc
// @Summary Execute payment contract
// @Description Execute payment for a contract invoice
// @Tags jobs
// @Produce json
// @Param invoiceID query string true "Invoice ID"
// @Success 200 {object} Job.ExecutePaymentRes "Successfully executed payment"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /job/payment/execute [post]
func (h *JobsHandler) ExecutePaymentContract(c *fiber.Ctx) error {
	invoiceID := c.Query("invoiceID")
	fmt.Println("1")
	res, err := h.job.ExecutePaymentContract(context.Background(), &Job.ExecutePaymentReq{InvoiceId: invoiceID})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println("2")

	return c.Status(int(res.Status)).JSON(fiber.Map{
		"orderID":     res.PaymentID,
		"merchantIDs": res.MerchantID,
	})
}

// GetPaymentContract godoc
// @Summary Get payment contract
// @Description Get details of a payment contract
// @Tags jobs
// @Produce html
// @Param invoiceID query string true "Invoice ID"
// @Success 200 {string} string "Payment contract details"
// @Router /jobs/payment [get]
func (h *JobsHandler) GetPaymentContract(c *fiber.Ctx) error {
	invoiceID := c.Query("invoiceID")
	fmt.Println(invoiceID)
	return c.Render("/home/jasim/GigForge-Freelance-Marketplace/API-Gateway/template/index.html", nil)
}

// CapturePaymentContract godoc
// @Summary Capture payment contract
// @Description Capture payment for a contract
// @Tags jobs
// @Produce json
// @Param paymentID query string true "Payment ID"
// @Param invoiceID query string true "Invoice ID"
// @Success 200 {object} Job.CapturePaymentRes "Successfully captured payment"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /jobs/payment/capture [post]
func (h *JobsHandler) CapturePaymentContract(c *fiber.Ctx) error {
	paymentID := c.Query("paymentID")
	invoiceID := c.Query("invoiceID")
	res, err := h.job.CapturePaymentContract(context.Background(), &Job.CapturePaymentReq{PaymentID: paymentID, InvoiceID: invoiceID})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res.UserName)
}

// Search godoc
// @Summary Search jobs
// @Description Search for jobs based on various criteria
// @Tags jobs
// @Produce json
// @Param q query string true "Query string"
// @Param t query string false "Pay type"
// @Param hourly_rate query string false "Hourly rate"
// @Param fixed_rate query string false "Fixed rate"
// @Param c query string false "Category"
// @Success 200 {object} Job.SearchJobsRes "Successfully retrieved search results"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /jobs/search [get]
func (h *JobsHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	PayType := c.Query("t")
	Hourly_rate := c.Query("hourly_rate")
	fixed_rate := c.Query("fixed_rate")
	category := c.Query("c")

	res, err := h.job.SearchJobs(context.Background(), &Job.SearchJobsReq{Query: query, Paytype: PayType, HourlyRate: Hourly_rate, FixedRate: fixed_rate, Category: category})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)

}

// GetAllContractsForClient godoc
// @Summary Get all contracts for a client
// @Description Get all contracts for a specific client
// @security ClientAccessToken
// @Tags contracts
// @Produce json
// @Param status query string false "Contract status"
// @Success 200 {object} Job.GetAllContractsForClientRes "Successfully retrieved contracts"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /client/contracts [get]
func (h *JobsHandler) GetAllContractsForClient(c *fiber.Ctx) error {
	userID := c.Locals("User_id").(int)
	Status := c.Query("status")
	contracts, err := h.job.GetAllContractsForClient(context.Background(), &Job.GetAllContractsForClientReq{UserId: int64(userID), Status: Status})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(contracts.Status)).JSON(contracts)
}

// GetOneContract godoc
// @Summary Get a contract
// @Description Get details of a specific contract
// @security ClientAccessToken
// @Tags contracts
// @Produce json
// @Param id path string true "Contract ID"
// @Success 200 {object} Job.GetOneContractForClientRes "Successfully retrieved contract"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /client/contracts/{id} [get]
func (h *JobsHandler) GetOneContract(c *fiber.Ctx) error {
	userID := c.Locals("User_id").(int)
	contractID := c.Params("id")
	contracts, err := h.job.GetOneContractForClient(context.Background(), &Job.GetOneContractForClientReq{UserId: int64(userID), ContractID: contractID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(contracts.Status)).JSON(contracts)
}

// GetAllInvoicesOfContract godoc
// @Summary Get all invoices of a contract
// @Description Get all invoices for a specific contract
// @security ClientAccessToken
// @Tags contracts
// @Produce json
// @Param contractID path string true "Contract ID"
// @Success 200 {object} Job.GetInvoiceContractRes "Successfully retrieved invoices"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /contracts/invoices/{contractID} [get]
func (h *JobsHandler) GetAllInvoicesOfContract(c *fiber.Ctx) error {
	userID := c.Locals("User_id").(int)
	cid := c.Params("contractID")
	contracts, err := h.job.GetInvoiceContract(context.Background(), &Job.GetInvoiceContractReq{UserID: int64(userID), ContractID: cid})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(contracts.Status)).JSON(contracts)
}

func (h *JobsHandler) AddContractAttachment(c *fiber.Ctx) error {
	var req req.AddContractAttachment
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}
	file, err := c.FormFile("attachment")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": Error})
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res,err:=h.job.AddAttachmentToContract(context.Background(),&Job.AddAttachmentReq{Attachment: fileBytes,Filename: file.Filename,ContractID: strconv.Itoa(req.ContractID),Description: req.Description})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetAttachments(c *fiber.Ctx) error{
	cid:=c.Params("contractID")
	res,err:=h.job.GetAttachments(context.Background(),&Job.GetAttachmentReq{ContractID: cid})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}