package handler

import (
	"context"
	"fmt"

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
		Title:           req.Title,
		Description:     req.Description,
		Category:        req.Category,
		Skills:          req.Skills,
		TimePeriod:      req.TimePeriod,
		FreelancerLevel: req.Level,
		Budget:          req.Budget,
		UserId:          user_id,
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

func (h *JobsHandler) ViewJobsForClient(c *fiber.Ctx) {

}

func (h *JobsHandler) EditJob(c *fiber.Ctx) error{
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
		Title:           req.Title,
		Description:     req.Description,
		Category:        req.Category,
		Skills:          req.Skills,
		TimePeriod:      req.TimePeriod,
		FreelancerLevel: req.Level,
		Budget:          req.Budget,
		UserId:          user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) ViewJobsForFreelancers(c *fiber.Ctx) {

}

func (h *JobsHandler) ViewProposalsForJob(c *fiber.Ctx) {

}

func (h *JobsHandler) SendOffer(c *fiber.Ctx) {

}

func (h *JobsHandler) AcceptOffer(c *fiber.Ctx) {

}
