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

// AdminLogin godoc
// @Summary Admin login
// @Description Authenticate an admin user with email and password
// @Tags admin
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login credentials"
// @Success 200 {object} auth.LoginRes "Successfully authenticated"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/login [post]
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

// AddSkill godoc
// @Summary Add a new skill
// @Description Add a new skill with the provided details
// @security AdminAccessToken
// @Tags admin
// @Accept json
// @Produce json
// @Param skill body req.AddSkills true "Skill details"
// @Success 200 {object} auth.AddSkillRes "Successfully added skill"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/skills [post]
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

// BlockUser godoc
// @Summary Block a user
// @Description Block a user by their ID
// @security AdminAccessToken
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} auth.BlockRes "User blocked successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/user/block [put]
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

// UnBlockUser godoc
// @Summary Unblock a user
// @Description Unblock a user by their ID
// @security AdminAccessToken
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} auth.BlockRes "User unblocked successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/user/unblock [put]
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

// AddCategory godoc
// @Summary Add a new category
// @Description Add a new category with the provided details
// @security AdminAccessToken
// @Tags admin
// @Accept json
// @Produce json
// @Param category body req.AddCategory true "Category details"
// @Success 200 {object} Job.AddCategoryRes "Successfully added category"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/category [post]
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

func (h *AdminHandler) AdminContractDashboard(c *fiber.Ctx) error{
	dashboard,err:=h.job.AdminContractDashboard(context.Background(),&Job.ACDReq{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.Status(int(dashboard.Status)).JSON(dashboard)
}