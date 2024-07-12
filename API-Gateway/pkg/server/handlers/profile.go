package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/notification"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct {
	profile auth.AuthServiceClient
	notification notification.NotificationServiceClient
}

func NewProfilehandler(profile auth.AuthServiceClient, notification notification.NotificationServiceClient) *ProfileHandler {
	return &ProfileHandler{profile: profile,notification: notification}
}

// AddEducationDetails adds education details for the user.
// @Summary Add education details
// @Description Add education details for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Education body req.Education true "Education Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/education [post]
func (h *ProfileHandler) AddEducationDetails(c *fiber.Ctx) error {
	var req req.Education
	user_id, _ := c.Locals("User_id").(int)
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
		UserId:       strconv.Itoa(user_id),
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

// UpdateEducation updates education details for the user.
// @Summary Update education details
// @Description Update education details for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param id path string true "Education ID"
// @Param Education body req.Education true "Education Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/education/{id} [patch]
func (h *ProfileHandler) UpdateEducation(c *fiber.Ctx) error {
	var req req.Education
	user_id := c.Locals("User_id").(int)
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
		UserId:       strconv.Itoa(user_id),
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

// DeleteEducation deletes education details for the user.
// @Summary Delete education details
// @Description Delete education details for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Produce json
// @Param id path string true "Education ID"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/education/{id} [delete]
func (h *ProfileHandler) DeleteEducation(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	e_id := c.Params("id")
	res, err := h.profile.DeleteEducation(context.Background(), &auth.DeleteEducationReq{
		UserId:      strconv.Itoa(user_id),
		EducationId: e_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// AddProfileDescription adds a profile description for the user.
// @Summary Add profile description
// @Description Add profile description for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Profile body req.Profile true "Profile Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/description [post]
func (h *ProfileHandler) AddProfileDescription(c *fiber.Ctx) error {
	var req req.Profile
	user_id := c.Locals("User_id").(int)
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
		UserId:      strconv.Itoa(user_id),
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// EditProfileDescription updates the profile description for the user.
// @Summary Update profile description
// @Description Update profile description for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Profile body req.Profile true "Profile Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/description [patch]
func (h *ProfileHandler) EditProfileDescription(c *fiber.Ctx) error {
	var req req.Profile
	user_id := c.Locals("User_id").(int)
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
		return c.Status(400).JSON(fiber.Map{"error":Error})
	}
	res, err := h.profile.UpdateProfileDescription(context.Background(), &auth.UPDReq{
		Title:       req.Title,
		Description: req.Description,
		HourlyRate:  req.Hourly_rate,
		UserId:      strconv.Itoa(user_id),
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// UpdateSkilltoProfile updates skills for the user profile.
// @Summary Update skills
// @Description Update skills for the user profile
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Skills body req.Skills true "Skills Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/skill [post]
func (h *ProfileHandler) UpdateSkilltoProfile(c *fiber.Ctx) error {
	var skill req.Skills
	user_id := c.Locals("User_id").(int)
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
		UserId: strconv.Itoa(user_id),
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}


// UpdateProfilePhoto updates the profile photo for the user.
// @Summary Update profile photo
// @Description Update profile photo for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept multipart/form-data
// @Produce json
// @Param profile-photo formData file true "Profile Photo"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/photo [put]
func (h *ProfileHandler) UpdateProfilePhoto(c *fiber.Ctx) error {
	userID := c.Locals("User_id").(int)
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
		UserId: strconv.Itoa(userID),
		Image:  imageData,
		Filename: file.Filename,
	})
	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(int(res.Status)).JSON(res)
}

// AddExperience adds an experience entry for the user.
// @Summary Add experience
// @Description Add experience for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Experience body req.Experience true "Experience Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/experience [post]
func (h *ProfileHandler) AddExperience(c *fiber.Ctx) error {
	var req req.Experience
	user_id := c.Locals("User_id").(int)
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
		UserId:      strconv.Itoa(user_id),
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// UpdateExperience updates an experience entry for the user.
// @Summary Update experience
// @Description Update experience for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param id path string true "Experience ID"
// @Param Experience body req.Experience true "Experience Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/experience/{id} [patch]
func (h *ProfileHandler) UpdateExperience(c *fiber.Ctx) error {
	var req req.Experience
	user_id := c.Locals("User_id").(int)
	exp_id := c.Params("id")
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON("error while parsing body.Check syntax")
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error":Error})
	}
	res, err := h.profile.UpdateExperience(context.Background(), &auth.ExpReq{
		Company:     req.Company,
		City:        req.City,
		Country:     req.Country,
		Title:       req.Title,
		FromDate:    req.FromDate,
		ToDate:      req.ToDate,
		Description: req.Description,
		UserId:      strconv.Itoa(user_id),
		ExpId:       exp_id,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// RemoveExperience deletes an experience entry for the user.
// @Summary Remove experience
// @Description Remove experience for the user
// @security FreelancerAccessToken
// @Tags Profile
// @Produce json
// @Param id path string true "Experience ID"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /freelancer/profile/experience/{id} [delete]
func (h *ProfileHandler) RemoveExperience(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	exp_id := c.Params("id")
	res, err := h.profile.DeleteExperience(context.Background(), &auth.DltExpReq{
		UserId:       strconv.Itoa(user_id),
		ExperienceId: exp_id,
	})
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// GetFreelancerProfile retrieves the profile of a freelancer.
// @Summary Get freelancer profile
// @Description Get the profile of a freelancer
// @security FreelancerAccessToken
// @Tags Profile
// @Produce json
// @Success 200 {object} res.CommonRes
// @Failure 403 {object} res.CommonRes
// @Router /freelancer/profile [get]
func (h *ProfileHandler) GetFreelancerProfile(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int)
	res, err := h.profile.GetProfile(context.Background(), &auth.GetProfileReq{UserId: strconv.Itoa(user_id)})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *ProfileHandler) OnboardFreelancersToPaypal(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(int)
	res, err := h.profile.OnboardFreelancerToPaypal(context.Background(), &auth.OnboardToPaypalReq{UserId: strconv.Itoa(user_id)})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}
	return c.Status(int(res.Status)).JSON(res)
}

// ReviewFreelancer adds a review for a freelancer.
// @Summary Add review for freelancer
// @Description Add a review for a freelancer
// @security ClientAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param Review body req.AddReview true "Review Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /client/review-freelancer [post]
func (h *ProfileHandler) ReviewFreelancer(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(int)
	var req req.AddReview

	if err:=c.BodyParser(&req);err != nil {
		return c.Status(400).JSON(fiber.Map{"error":"error in parsing body. enter fields correctly"})
	}
	res,err:=h.profile.ReviewFreelancer(context.Background(),&auth.ReviewFlancerReq{Review: req.Review,FreelancerId: int32(req.Freelancer_id),Rating: int32(req.Rating),ClientId: int64(user_id)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetTalents retrieves a list of talents.
// @Summary Get talents
// @Description Get a list of talents
// @Tags Profile
// @Produce json
// @Param q query string false "Query"
// @Param exp query string false "Experience"
// @Success 200 {object} res.CommonRes
// @Failure 500 {object} res.CommonRes
// @Router /talents [get]
func (h *ProfileHandler) GetTalents(c *fiber.Ctx) error{
	query:=c.Query("q")
	exp:=c.Query("exp")
	talents,err:=h.profile.GetFreelancers(context.Background(),&auth.GetTalentReq{Query: query,Exp: exp})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(talents.Status)).JSON(talents)
}

// AddPaymentEmailPaypal adds a PayPal payment email for the user.
// @Summary Add PayPal payment email
// @Description Add a PayPal payment email for the user
// @Param Authorization header string true "Authorization"
// @Tags Profile
// @Accept json
// @Produce json
// @Param Payment body req.AddPayment true "Payment Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /payment/add [post]
func (h *ProfileHandler) AddPaymentEmailPaypal(c *fiber.Ctx) error{
	user_id:=c.Locals("User_id").(int)
	var req req.AddPayment
	if err:=c.BodyParser(&req);err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"error while validating body.Enter correctly"})
	}

	res,err:=h.profile.AddPaymentEmail(context.Background(),&auth.AddPaymentEmailReq{UserId: strconv.Itoa(user_id),Email: req.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}



// UpdateCompanyDetails updates the company details for the user.
// @Summary Update company details
// @Description Update the company details for the user
// @security ClientAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param CompanyDetails body req.UpdateCompanyDetails true "Company Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /client/profile/company-details [put]
func (h *ProfileHandler) UpdateCompanyDetails(c *fiber.Ctx) error{
	var req req.UpdateCompanyDetails
	user_id:=c.Locals("User_id").(int)
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

// GetClientProfile retrieves the profile of a client.
// @Summary Get client profile
// @Description Get the profile of a client
// @security ClientAccessToken
// @Tags Profile
// @Produce json
// @Success 200 {object} res.CommonRes
// @Failure 403 {object} res.CommonRes
// @Router /client/profile [get]
func (h *ProfileHandler) GetClientProfile(c *fiber.Ctx) error{
	user_id := c.Locals("User_id").(int)
	
	res, err := h.profile.GetProfileClient(context.Background(), &auth.ClientProfileReq{UserId:int32(user_id)})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

// UpdateCompanyContacts updates the company contact details for the user.
// @Summary Update company contact details
// @Description Update the company contact details for the user
// @security ClientAccessToken
// @Tags Profile
// @Accept json
// @Produce json
// @Param CompanyContact body req.UpdateCompanyContact true "Company Contact Details"
// @Success 200 {object} res.CommonRes
// @Failure 400 {object} res.CommonRes
// @Router /client/profile/company-contacts [put]
func (h *ProfileHandler) UpdateCompanyContacts(c *fiber.Ctx) error{
	var req req.UpdateCompanyContact
	user_id:=c.Locals("User_id").(int)
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

// GetFreelancerReviews retrieves reviews for a freelancer.
// @Summary Get freelancer reviews
// @Description Get reviews for a freelancer
// @Tags Profile
// @Produce json
// @Param freelancer_id path string true "Freelancer ID"
// @Success 200 {object} res.CommonRes
// @Failure 500 {object} res.CommonRes
// @Router /reviews/{freelancer_id} [get]
func (h *ProfileHandler) GetFreelancerReviews(c *fiber.Ctx) error{
	fid:=c.Params("freelancer_id")
	res,err:=h.profile.GetFreelancerReviews(context.Background(),&auth.GetReviewReq{UserID:fid })
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

// GetNotifications retrieves notifications for the user.
// @Summary Get user notifications
// @Description Get notifications for the user
// @security Authorization
// @Tags Profile
// @Produce json
// @Param UserId path int true "User ID"
// @Success 200 {object} res.CommonRes
// @Failure 500 {object} res.CommonRes
// @Router /notifications [get]
func (h *ProfileHandler) GetNotifications(c *fiber.Ctx) error{
	userID,exist:=c.Locals("UserId").(int)
	if !exist{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "login to see notifications"})
	}
	notifications,err:=h.notification.GetNotification(context.Background(),&notification.GNReq{UserId: int32(userID)})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(int(notifications.Status)).JSON(notifications)
}