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

func (h *ProjectHandler) RemoveProject(c *fiber.Ctx) error {
	prjt_id:=c.Params("id")
	user_id:=c.Locals("User_id").(string)

	res,err:=h.project.RemoveProject(context.Background(),&project.RemProjectReq{UserId: user_id,ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProjectHandler) ListProjects(c *fiber.Ctx) error {
	res,err:=h.project.ListProjects(context.Background(),&project.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProjectHandler) ListProjectWithID(c *fiber.Ctx) error {
	prjt_id:=c.Params("id")
	res,err:=h.project.ListOneProject(context.Background(),&project.ListOneProjectReq{ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProjectHandler) ListMyProjects(c *fiber.Ctx) error {
	user_id:=c.Locals("User_id").(string)
	res,err:=h.project.ListMyProjects(context.Background(),&project.ListMyProjectReq{UserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})

	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProjectHandler) BuyProject(c *fiber.Ctx) error {
	user_id:=c.Locals("User_id").(int64)
	user:=strconv.Itoa(int(user_id))
	prjt_id:=c.Params("id")
	res,err:=h.project.BuyProject(context.Background(),&project.BuyProjectReq{UserId: user,ProjectId: prjt_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// func (h *ProjectHandler) SearchProject(c *fiber.Ctx) error{
// 	query:=c.Params("q")
// }

// func (h *ProjectHandler) ExecutePaymentForProject(c *fiber.Ctx) error {
	
// }
