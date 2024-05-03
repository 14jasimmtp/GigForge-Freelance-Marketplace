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

func (h *JobsHandler) GetMyJobs(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int64)
	id := strconv.Itoa(int(user_id))
	res, err := h.job.GetMyJobs(context.Background(), &Job.GetMyJobsReq{UserId: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// func (h *JobsHandler) GetMyProposals(c *fiber.Ctx) error {
// 	user_id := c.Locals("User_id").(string)
// 	res, err := h.job.GetMyProposals(context.Background(), &Job.ProposalReq{})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
// 	}
// 	return c.Status(int(res.Status)).JSON(res)
// }

// func (h *JobsHandler) GetProposals(c *fiber.Ctx)  error{

// }

// func (h *JobsHandler) GetJob(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{})
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.Status(res.Status).JSON(res)
// }

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

func (h *JobsHandler) AcceptOffer(c *fiber.Ctx) error{
	user_id := c.Locals("User_id").(string)
	of_id:=c.Params("offer_id")
	// if err := c.BodyParser(&req); err != nil {
	// 	return c.Status(400).JSON(
	// 		res.CommonRes{
	// 			Status:  "failed",
	// 			Message: "Error validating request body",
	// 			Error:   err.Error(),
	// 			Body:    nil,
	// 		},
	// 	)
	// }

	// Error, err := validation.Validation(req)
	// if err != nil {
	// 	return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	// }
	res,err:=h.job.AcceptOffer(context.Background(),&Job.AcceptOfferReq{UserId: user_id,OfferID: of_id})
	if err != nil {
		print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"something went wrong"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJobs(c *fiber.Ctx) error{
	res, err := h.job.GetJobs(context.Background(), &Job.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJob(c *fiber.Ctx) error{
	job_id:=c.Params("id")
	println(job_id)
	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) SendInvoice(c *fiber.Ctx) error{
	var req req.SendInvoice
	user_id := c.Locals("User_id").(string)
	res,err:=h.job.SendWeeklyInvoice(context.Background(),&Job.InvoiceReq{ContractID:int32(req.ContractId),TotalHourWorked: float32(req.TotalHoursWorked),SuserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err}) 
	}
	return c.Status(int(res.Status)).JSON(res)
}