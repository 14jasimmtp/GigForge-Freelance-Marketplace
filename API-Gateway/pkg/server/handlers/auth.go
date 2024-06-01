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

// Login godoc
// @Summary User login
// @Description Authenticate a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body req.LoginRequest true "Login credentials"
// @Success 200 {object} auth.UserLoginRes "Successfully authenticated"
// @Failure 400 {object} res.CommonRes "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/login [post]
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
		return c.Status(400).JSON(fiber.Map{"error":Error})
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

// Signup godoc
// @Summary User signup
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param role query string true "User role"
// @Param user body req.SignupRequest true "User details"
// @Success 200 {object} auth.UserSignupRes "Successfully signed up"
// @Failure 400 {object} res.CommonRes "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/signup [post]
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
		return c.Status(400).JSON(fiber.Map{"error": Error})
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

// ForgotPassword godoc
// @Summary Forgot password
// @Description Request to reset user password
// @Tags auth
// @Accept json
// @Produce json
// @Param req body req.ForgotPassword true "Request details"
// @Success 200 {object} auth.FPres "Password reset request successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/forgot-password [post]
func (h *Handler) ForgotPassword(c *fiber.Ctx) error {
	var req req.ForgotPassword

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(err)
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": Error})
	}
	res, err := h.auth.ForgotPassword(context.Background(), &auth.FPreq{
		Email: req.Email,
	})
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)

}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset user password using OTP and token
// @Tags auth
// @Accept json
// @Produce json
// @Param req body req.ResetPassword true "Request details"
// @Param Authorization header string true "Authentication token"
// @Success 200 {object} auth.RPres "Password reset successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/reset-password [post]
func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var req req.ResetPassword

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": Error})
	}
	token:=c.Get("Authorization")

	res, err := h.auth.ResetPassword(context.Background(), &auth.RPreq{
		Password: req.NewPassword,
		OTP:      req.OTP,
		Token: token,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)
}

// Verify godoc
// @Summary Verify OTP
// @Description Verify OTP for user authentication
// @Tags auth
// @Accept json
// @Produce json
// @Param req body req.Verify true "Verification details"
// @Param Authorization header string true "Authentication token"
// @Success 200 {object} auth.VerifyRes "OTP verification successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/verify [post]
func (h *Handler) Verify(c *fiber.Ctx) error {
	var req req.Verify

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": Error})
	}

	token := c.Get("Authorization")
	fmt.Println("token",token)

	res, err := h.auth.Verify(context.Background(), &auth.VerifyReq{
		OTP:   req.OTP,
		Token: token,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

