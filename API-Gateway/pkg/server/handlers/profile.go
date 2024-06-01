package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	profile auth.AuthServiceClient
}

func NewProfilehandler(profile auth.AuthServiceClient) *ProfileHandler {
	return &ProfileHandler{profile: profile}
}

//Freelancer profile

func (h *ProfileHandler) AddEducationDetails(c *fiber.Ctx) error {
	var req req.Education
	user_id, _ := c.Locals("User_id").(string)
	fmt.Println("user", user_id)
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
		return c.Status(400).JSON(fiber.Map{"Error":Error})
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
	user_id := c.Locals("User_id").(string)
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
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
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
	user_id := c.Locals("User_id").(string)
	e_id := c.Params("id")
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
		return c.Status(400).JSON(Error)
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

func (h *ProfileHandler) UpdateSkilltoProfile(c *fiber.Ctx) error {
	var skill req.Skills
	user_id := c.Locals("User_id").(string)
	if err := c.BodyParser(&skill); err != nil {
		return c.Status(400).JSON(
			res.CommonRes{
				Status:  "failed",
				Message: "Error validating request body",
				Error:   err.Error(),
				Body:    nil,
			},
		)
	}
	Error, err := validation.Validation(skill)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}
	res, err := h.profile.EditSkill(context.Background(), &auth.EditSkillReq{
		Skills: skill.Skills,
		UserId: user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) UpdateProfilePhoto(c *fiber.Ctx) error {
	userID := c.Locals("User_id").(string)
	file, err := c.FormFile("profile-photo")
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error":"file not found"})
	}
	if !validation.IsJPEG(file){
		return c.Status(404).JSON(fiber.Map{"error":"file should jpg format"})
	}

	fileContent, err := file.Open()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	defer fileContent.Close()

	imageData, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(400).JSON(err)
	}
	res, err := h.profile.UpdateProfilePhoto(context.Background(), &auth.PhotoReq{
		UserId: userID,
		Image:  imageData,
	})
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) AddExperience(c *fiber.Ctx) error {
	var req req.Experience
	user_id := c.Locals("User_id").(string)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON("error while parsing body.Check syntax")
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}
	res, err := h.profile.AddExperience(context.Background(), &auth.ExpReq{
		Company:     req.Company,
		City:        req.City,
		Country:     req.Country,
		Title:       req.Title,
		FromDate:    req.FromDate,
		ToDate:      req.ToDate,
		Description: req.Description,
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) UpdateExperience(c *fiber.Ctx) error {
	var req req.Experience
	user_id := c.Locals("User_id").(string)
	exp_id := c.Params("id")
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON("error while parsing body.Check syntax")
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fmt.Sprintf(`{"error": %v}`, Error))
	}
	res, err := h.profile.UpdateExperience(context.Background(), &auth.ExpReq{
		Company:     req.Company,
		City:        req.City,
		Country:     req.Country,
		Title:       req.Title,
		FromDate:    req.FromDate,
		ToDate:      req.ToDate,
		Description: req.Description,
		UserId:      user_id,
		ExpId:       exp_id,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) RemoveExperience(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	exp_id := c.Params(":id")
	res, err := h.profile.DeleteExperience(context.Background(), &auth.DltExpReq{
		UserId:       user_id,
		ExperienceId: exp_id,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) GetFreelancerProfile(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	res, err := h.profile.GetProfile(context.Background(), &auth.GetProfileReq{UserId: user_id})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) OnboardFreelancersToPaypal(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(string)
	res, err := h.profile.OnboardFreelancerToPaypal(context.Background(), &auth.OnboardToPaypalReq{UserId: user_id})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) ReviewFreelancer(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(int64)
	var req req.AddReview

	if err:=c.BodyParser(&req);err != nil {
		return c.Status(400).JSON(fiber.Map{"error":"error in parsing body. enter fields correctly"})
	}
	res,err:=h.profile.ReviewFreelancer(context.Background(),&auth.ReviewFlancerReq{Review: req.Review,FreelancerId: int32(req.Freelancer_id),Rating: int32(req.Rating),ClientId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) GetTalents(c *fiber.Ctx) error{
	query:=c.Query("q")
	exp:=c.Query("exp")
	talents,err:=h.profile.GetFreelancers(context.Background(),&auth.GetTalentReq{Query: query,Exp: exp})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(talents.Status)).JSON(talents)
}

func (h *ProfileHandler) AddPaymentEmailPaypal(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(string)
	var req req.AddPayment
	if err:=c.BodyParser(&req);err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"error while validating body.Enter correctly"})
	}

	res,err:=h.profile.AddPaymentEmail(context.Background(),&auth.AddPaymentEmailReq{UserId: user_id,Email: req.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// func (h *ProfileHandler) GetFreelancerReviews(c *fiber.Ctx) error{

// }

func (h *ProfileHandler) UpdateCompanyDetails(c *fiber.Ctx) error{
	var req req.UpdateCompanyDetails
	user_id:=c.Locals("User_id").(int64)
	if err := c.BodyParser(&req); err != nil{
		return c.Status(400).JSON("error while parsing body.Check syntax")
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error":Error})
	}
	res,err:=h.profile.UpdateCompanyDetails(context.Background(),&auth.UpdCompDtlReq{
		CompanyName: req.CompanyName,
		Website: req.Website,
		NumberOfEmployees: int32(req.NoOfEmployees),
		Tagline: req.Tagline,
		Industry: req.Industry,
		UserId: int32(user_id),
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) GetClientProfile(c *fiber.Ctx) error{
	user_id := c.Locals("User_id").(int64)
	
	res, err := h.profile.GetProfileClient(context.Background(), &auth.ClientProfileReq{UserId:int32(user_id)})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) UpdateCompanyContacts(c *fiber.Ctx) error{
	var req req.UpdateCompanyContact
	user_id:=c.Locals("User_id").(int64)
	if err := c.BodyParser(&req); err != nil{
		return c.Status(400).JSON("error while parsing body.Check syntax")
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error":Error})
	}
	res,err:=h.profile.UpdateCompanyContact(context.Background(),&auth.UpdCompContReq{
		UserId: int64(user_id),
		OwnerName: req.OwnerName,
		Phone: req.Phone,
		Address: &auth.Address{
			Country: req.Address.Country,
			State: req.Address.State,
			District: req.Address.District,
			City: req.Address.City,
			Pincode: req.Address.PinCode,
		},
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) GetFreelancerReviews(c *fiber.Ctx) error{
	fid:=c.Params("Fid")
	res,err:=h.profile.GetFreelancerReviews(context.Background(),&auth.GetReviewReq{UserID:fid })
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}


// func (h *ProfileHandler) GetPaymentHistory(c *fiber.Ctx) error{
	
// }