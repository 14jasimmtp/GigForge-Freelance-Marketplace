package handler

import (
	"context"
	"fmt"
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

func (h *JobsHandler) PostJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int64)
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
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)

}

func (h *JobsHandler) SendProposal(c *fiber.Ctx) error {
	var req req.Proposal

	job_id := c.Query("jobID")
	user_id := c.Locals("User_id").(string)
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
		UserId:      user_id,
		JobId:       job_id,
		// Attachments: attachments,
	})

	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetProposalsOfJob(c *fiber.Ctx) error {
	client_id := strconv.Itoa(int(c.Locals("User_id").(int64)))
	job_id := c.Params("id")

	res, err := h.job.GetJobProposals(context.Background(), &Job.GJPReq{UserId: client_id, JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetMyJobs(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int64)
	id := strconv.Itoa(int(user_id))
	res, err := h.job.GetMyJobs(context.Background(), &Job.GetMyJobsReq{UserId: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetCategories(c *fiber.Ctx) error{
	query:=c.Query("q")
	category, err := h.job.GetCategory(context.Background(),&Job.GetCategoryReq{Query: query})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error":err.Error()})
	}
	return c.Status(int(category.Status)).JSON(category)
}

func (h *JobsHandler) GetJobProposals(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	user_id := c.Locals("User_id").(string)
	res, err := h.job.GetJobProposals(context.Background(), &Job.GJPReq{JobId: jobID, UserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) EditJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int64)
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
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) SendOffer(c *fiber.Ctx) error {
	var req req.SendOffer
	user_id := c.Locals("User_id").(int64)

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

func (h *JobsHandler) AcceptOffer(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	of_id := c.Params("offer_id")

	res, err := h.job.AcceptOffer(context.Background(), &Job.AcceptOfferReq{UserId: user_id, OfferID: of_id})
	if err != nil {
		print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJobs(c *fiber.Ctx) error {
	res, err := h.job.GetJobs(context.Background(), &Job.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJob(c *fiber.Ctx) error {
	job_id := c.Params("id")
	println(job_id)
	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetAllJobOffersForFreelancer(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	status:=c.Query("status")
	res, err := h.job.GetJobOffersForFreelancer(context.Background(), &Job.GetJobOfferForFreelancerReq{UserId: user_id,Status: status})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) SendInvoice(c *fiber.Ctx) error {
	var req req.SendInvoice
	user_id := c.Locals("User_id").(string)
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
	res, err := h.job.SendWeeklyInvoice(context.Background(), &Job.InvoiceReq{ContractID: int32(req.ContractId), TotalHourWorked: float32(req.TotalHoursWorked), SuserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

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

func (h *JobsHandler) GetPaymentContract(c *fiber.Ctx) error{
	invoiceID:=c.Query("invoiceID")
	fmt.Println(invoiceID)
	return c.Render("/home/jasim/GigForge-Freelance-Marketplace/API-Gateway/template/index.html",nil)
}

func (h *JobsHandler) CapturePaymentContract(c *fiber.Ctx) error{
	paymentID := c.Query("paymentID")
	invoiceID := c.Query("invoiceID")
	res, err := h.job.CapturePaymentContract(context.Background(), &Job.CapturePaymentReq{PaymentID: paymentID,InvoiceID: invoiceID})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res.UserName)
}

// func (h *JobsHandler) CloseJobPost() {

// }

// func (h *JobsHandler) GetContractDetails(c *fiber.Ctx){
// 	job_id:=c.Params("id")
// 	println(job_id)
// 	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
// 	}
// 	return c.Status(int(res.Status)).JSON(res)
// }

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

//contracts

// func (h *JobsHandler) GetAllContractsForClient(c *fiber.Ctx) error {
// 	user_id := c.Locals("User_id").(int64)
// 	res, err := h.job.GetAllContractsForClient(context.Background(), &Job.GetAllContractsForClientReq{UserId: user_id})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(int(res.Status)).JSON(res)
// }

// func (h *JobsHandler) EndContract(c *fiber.Ctx) error {

// }

func (h *JobsHandler) GetPaymentForContractWithInvoiceID(c *fiber.Ctx){
	
}
