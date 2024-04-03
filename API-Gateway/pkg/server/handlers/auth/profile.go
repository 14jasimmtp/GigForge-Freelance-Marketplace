package handler

import (
	"context"
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	profile auth.AuthServiceClient
}

func NewProfilehandler(profile auth.AuthServiceClient) *ProfileHandler {
	return &ProfileHandler{profile: profile}
}

func (h *ProfileHandler) AddEducationDetails(c *fiber.Ctx) error {
	var req req.Education
	user_id,_ := c.Locals("User_id").(string)
	fmt.Println("user",user_id)
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

	res, err := h.profile.AddEducation(context.Background(), &auth.AddEducationReq{
		School:       req.School,
		UserId:       user_id,
		Course:       req.Course,
		Date_Started: req.Date_Started,
		Date_Ended:   req.Date_Ended,
		AreaOfStudy:  req.Area_Of_Study,
		Description:  req.Description,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) UpdateEducation(c *fiber.Ctx) error {
	var req req.Education
	user_id := c.Get("user_id")
	e_id := c.Params("id")
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

	res, err := h.profile.UpdateEducation(context.Background(), &auth.UpdateEducationReq{
		EducationId:  e_id,
		School:       req.School,
		UserId:       user_id,
		Course:       req.Course,
		Date_Started: req.Date_Started,
		Date_Ended:   req.Date_Ended,
		AreaOfStudy:  req.Area_Of_Study,
		Description:  req.Description,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) DeleteEducation(c *fiber.Ctx) error {
	var req req.Education
	user_id := c.Get("user_id")
	e_id := c.Params("id")
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

	res, err := h.profile.DeleteEducation(context.Background(), &auth.DeleteEducationReq{
		UserId:      user_id,
		EducationId: e_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) AddProfileDescription(c *fiber.Ctx) error {
	var req req.Profile
	user_id := c.Get("user_id")
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

	res, err := h.profile.AddProfileDescription(context.Background(), &auth.APDReq{
		Title:       req.Title,
		Description: req.Description,
		HourlyRate:  req.Hourly_rate,
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) EditProfileDescription(c *fiber.Ctx) error {
	var req req.Profile
	user_id := c.Get("user_id")
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

	res, err := h.profile.UpdateProfileDescription(context.Background(), &auth.UPDReq{
		Title:       req.Title,
		Description: req.Description,
		HourlyRate:  req.Hourly_rate,
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error{

// }

// func (h *ProfileHandler) GetProfileClient(c *fiber.Ctx) error{

// }
