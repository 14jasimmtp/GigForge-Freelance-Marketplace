package handler

import (
	"context"
	"fmt"

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

func (h *ProjectHandler) AddProject(c *fiber.Ctx) error {
	var req req.AddProject

	if err:=c.BodyParser(&req);err != nil {
		return c.JSON(fiber.Map{"error": "error parsing datas","message":err.Error()})
	}

	Errors, err := validation.Validation(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": Errors})
	}
	grpcReq := &project.AddProjectReq{}
	copier.Copy(grpcReq, &req)
	fmt.Println(grpcReq)
	
	res, err := h.project.AddProject(context.Background(), grpcReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) EditProject(c *fiber.Ctx) error {
	return nil
}

func (h *ProfileHandler) RemoveProject(c *fiber.Ctx) error {
	return nil
}

func (h *ProfileHandler) BuyProduct(c *fiber.Ctx) error {
	return nil
}
