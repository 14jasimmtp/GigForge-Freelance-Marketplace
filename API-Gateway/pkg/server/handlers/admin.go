package handler

import (
	"context"
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/Job"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	auth auth.AuthServiceClient
	job  Job.JobServiceClient
}

func NewAdminHandler(job Job.JobServiceClient, auth auth.AuthServiceClient) AdminHandler {
	return AdminHandler{auth: auth,job: job}
}

func (h *AdminHandler) AdminLogin(ctx *fiber.Ctx) error {
	var login req.LoginRequest

	if ctx.BodyParser(&login) != nil {
		return ctx.Status(400).JSON(`"error":"bodies not passed correctly"`)
	}

	res, err := h.auth.AdminLogin(context.Background(), &auth.LoginReq{
		Email:    login.Email,
		Password: login.Password,
	})
	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(`"Error" : "rpc error occured"`)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func (h *AdminHandler) AddSkill(ctx *fiber.Ctx) error {
	var skill req.AddSkills

	if ctx.BodyParser(&skill) != nil {
		return ctx.Status(400).JSON(`"error":"bodies not passed correctly"`)
	}

	res, err := h.auth.AddSkill(context.Background(), &auth.AddSkillReq{
		Skill:       skill.Skill,
		Description: skill.Description,
	})
	if err != nil {
		return ctx.Status(500).JSON(`"Error" : "rpc error occured"`)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func (h *AdminHandler) BlockUser(ctx *fiber.Ctx) error {
	user_id := ctx.Query("id")

	res, err := h.auth.BlockUser(context.Background(), &auth.BlockReq{
		UserId: user_id,
	})
	if err != nil {
		fmt.Println(err)

		return ctx.Status(500).JSON(`"Error" : "rpc error occured"`)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func (h *AdminHandler) UnBlockUser(ctx *fiber.Ctx) error {
	user_id := ctx.Query("id")

	res, err := h.auth.UnBlockUser(context.Background(), &auth.BlockReq{
		UserId: user_id,
	})
	if err != nil {
		return ctx.Status(500).JSON(`"Error" : "rpc error occured"`)
	}

	return ctx.Status(int(res.Status)).JSON(res)
}

func (h *AdminHandler) AddCategory(c *fiber.Ctx) error {
	var req req.AddCategory

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error while parsing the body"})
	}
	fmt.Println("1")
	res, err := h.job.AddCategory(context.Background(), &Job.AddCategoryReq{Category: req.Category})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)
}
