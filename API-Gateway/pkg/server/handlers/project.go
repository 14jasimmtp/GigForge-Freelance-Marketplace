package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/project"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type ProjectHandler struct {
	project project.ProjectServiceClient
}

func NewProjectHandler(p project.ProjectServiceClient) *ProjectHandler {
	return &ProjectHandler{project: p}
}

// AddSingleProject godoc
// @Summary Add a single project
// @Description Create a new project listing
// @security FreelancerAccessToken
// @Tags projects
// @Accept json
// @Produce json
// @Param project body req.AddSingleProject true "Project details"
// @Success 200 {object} project.AddSingleProjectRes "Successfully added project"
// @Failure 400 {object} map[string]string "Error parsing request body"
// @Failure 401 {object} map[string]interface{} "Validation errors"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/add [post]
func (h *ProjectHandler) AddSingleProject(c *fiber.Ctx) error {
	var req req.AddSingleProject
	user_id:=c.Locals("User_id").(string)
	if err:=c.BodyParser(&req);err != nil {
		return c.JSON(fiber.Map{"error": "error parsing datas","message":err.Error()})
	}

	Errors, err := validation.Validation(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": Errors})
	}
	grpcReq := &project.AddSingleProjectReq{}
	copier.Copy(grpcReq, &req)
	grpcReq.UserId=user_id

	res, err := h.project.AddProject(context.Background(), grpcReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProjectHandler) AddTieredProject(c *fiber.Ctx) error {
	var req req.AddSingleProject
	// user_id:=c.Locals("User_id").(string)
	if err:=c.BodyParser(&req);err != nil {
		return c.JSON(fiber.Map{"error": "error parsing datas","message":err.Error()})
	}

	Errors, err := validation.Validation(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": Errors})
	}
	grpcReq := &project.AddSingleProjectReq{}
	copier.Copy(grpcReq, &req)
	fmt.Println(grpcReq)

	res, err := h.project.AddProject(context.Background(), grpcReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)}

// EditProject godoc
// @Summary Edit a project
// @Description Edit an existing project listing
// @security FreelancerAccessToken
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body req.AddSingleProject true "Project details"
// @Success 200 {object} project.EditSingleProjectRes "Successfully edited project"
// @Failure 400 {object} map[string]string "Error parsing request body"
// @Failure 401 {object} map[string]interface{} "Validation errors"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/edit/{id} [put]
func (h *ProjectHandler) EditProject(c *fiber.Ctx) error {
	var req req.AddSingleProject
	prjt_id:=c.Params("id")
	user_id:=c.Locals("User_id").(string)
	if err:=c.BodyParser(&req);err != nil {
		return c.JSON(fiber.Map{"error": "error parsing datas","message":err.Error()})
	}

	Errors, err := validation.Validation(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": Errors})
	}

	res, err := h.project.EditProject(context.Background(), &project.EditSingleProjectReq{
		Title: req.Title,
		Description: req.Description,
		Category: int32(req.Category),
		Price: req.Price,
		DeliveryDays: req.DeliveryDays,
		NumberOfRevisions: req.NumberOfRevisions,
		ProjectId: prjt_id,
		UserId: user_id,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)}

// RemoveProject godoc
// @Summary Remove a project
// @Description Remove an existing project listing
// @security FreelancerAccessToken
// @Tags projects
// @Param id path string true "Project ID"
// @Success 200 {object} project.RemProjectRes "Successfully removed project"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/remove/{id} [delete]
func (h *ProjectHandler) RemoveProject(c *fiber.Ctx) error {
	prjt_id:=c.Params("id")
	user_id:=c.Locals("User_id").(string)

	res,err:=h.project.RemoveProject(context.Background(),&project.RemProjectReq{UserId: user_id,ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// ListProjects godoc
// @Summary List all projects
// @Description Get a list of all projects
// @Tags projects
// @Produce json
// @Success 200 {object} project.ListProjectsRes "Successfully retrieved projects"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects [get]
func (h *ProjectHandler) ListProjects(c *fiber.Ctx) error {
	res,err:=h.project.ListProjects(context.Background(),&project.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}
	return c.Status(int(res.Status)).JSON(res)
}

// ListProjectWithID godoc
// @Summary Get project by ID
// @Description Get details of a specific project
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} project.ListOneProjectRes "Successfully retrieved project"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/{id} [get]
func (h *ProjectHandler) ListProjectWithID(c *fiber.Ctx) error {
	prjt_id:=c.Params("id")
	res,err:=h.project.ListOneProject(context.Background(),&project.ListOneProjectReq{ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// ListMyProjects godoc
// @Summary List my projects
// @Description Get a list of projects posted by the authenticated user
// @security FreelancerAccessToken
// @Tags projects
// @Produce json
// @Success 200 {object} project.ListMyProjectRes "Successfully retrieved projects"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/my [get]
func (h *ProjectHandler) ListMyProjects(c *fiber.Ctx) error {
	user_id:=c.Locals("User_id").(string)
	res,err:=h.project.ListMyProjects(context.Background(),&project.ListMyProjectReq{UserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}
	return c.Status(int(res.Status)).JSON(res)
}

// BuyProject godoc
// @Summary Buy a project
// @Description Buy a specific project
// @security ClientAccessToken
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} project.BuyProjectRes "Successfully bought project"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/buy/{id} [post]
func (h *ProjectHandler) BuyProject(c *fiber.Ctx) error {
	user_id:=c.Locals("User_id").(int64)
	user:=strconv.Itoa(int(user_id))
	prjt_id:=c.Params("id")
	res,err:=h.project.OrderProject(context.Background(),&project.BuyProjectReq{UserId: user,ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// ExecutePaymentProject godoc
// @Summary Execute project payment
// @Description Execute payment for a project order
// @Tags projects
// @Produce json
// @Param orderID query string true "Order ID"
// @Success 200 {object} project.ExecutePaymentRes "Successfully executed payment"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/payment/execute [post]
func (h *ProjectHandler) ExecutePaymentProject(c *fiber.Ctx) error {
	orderID := c.Query("orderID")
	fmt.Println(orderID)
	res, err := h.project.ExecutePaymentProject(context.Background(), &project.ExecutePaymentReq{OrderID: orderID})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(int(res.Status)).JSON(fiber.Map{
        "orderID":     res.PaymentID,
        "merchantIDs": res.MerchantID,
    })
}

// GetPaymentProject godoc
// @Summary Get project payment details
// @Description Get details of a project payment
// @Tags projects
// @Produce html
// @Param orderID query string true "Order ID"
// @Success 200 {string} string "Payment project details"
// @Router /projects/payment [get]
func (h *ProjectHandler) GetPaymentProject(c *fiber.Ctx) error{
	c.Query("orderID")
	return c.Render("/home/jasim/GigForge-Freelance-Marketplace/API-Gateway/template/projectpay.html",nil)
}

// CapturePaymentProject godoc
// @Summary Capture project payment
// @Description Capture payment for a project order
// @Tags projects
// @Produce json
// @Param paymentID query string true "Payment ID"
// @Param orderID query string true "Order ID"
// @Success 200 {object} project.CapturePaymentRes "Successfully captured payment"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /projects/payment/capture [post]
func (h *ProjectHandler) CapturePaymentProject(c *fiber.Ctx) error{
	paymentID := c.Query("paymentID")
	orderID := c.Query("orderID")
	res, err := h.project.CapturePaymentProject(context.Background(), &project.CapturePaymentReq{PaymentID: paymentID,OrderID: orderID})
	if err != nil {
		return c.Status(int(res.Status)).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res.UserName)
}
