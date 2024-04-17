package handler

import (
	"context"
	"fmt"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	auth auth.AuthServiceClient
}

func NewAuthHandler(auth auth.AuthServiceClient) *Handler {
	return &Handler{auth: auth}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var Login req.LoginRequest

	if err := c.BodyParser(&Login); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error while parsing body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}
	Error, err := validation.Validation(Login)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	res, err := h.auth.Login(context.Background(), &auth.UserLoginReq{
		Email:    Login.Email,
		Password: Login.Password,
	})

	if err != nil {
		return c.JSON(err)
	}

	return c.Status(int(res.Status)).JSON(res)

}

func (h *Handler) Signup(c *fiber.Ctx) error {
	var user req.SignupRequest

	role := c.Query("role")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}

	Error, err := validation.Validation(user)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	res, err := h.auth.Signup(context.Background(), &auth.UserSignupReq{
		Email:     user.Email,
		Password:  user.Password,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Phone:     user.Phone,
		Country:   user.Country,
		Role:      role,
	})

	if err != nil {
		return c.JSON(err)
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *Handler) ForgotPassword(c *fiber.Ctx) error {
	var req req.ForgotPassword

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(err)
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}
	res, err := h.auth.ForgotPassword(context.Background(), &auth.FPreq{
		Email: req.Email,
	})
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)

}

func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var req req.ResetPassword

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}
	token:=c.Get("Authorization")

	res, err := h.auth.ResetPassword(context.Background(), &auth.RPreq{
		Password: req.NewPassword,
		OTP:      req.OTP,
		Token: token,
	})
	if err != nil {
		return c.Status(500).JSON(err, "rpc error")
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *Handler) Verify(c *fiber.Ctx) error {
	var req req.Verify

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}

	token := c.Get("Authorization")

	res, err := h.auth.Verify(context.Background(), &auth.VerifyReq{
		OTP:   req.OTP,
		Token: token,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(400).JSON(res)
}
