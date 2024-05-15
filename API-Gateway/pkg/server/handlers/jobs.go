package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pb/Job"
	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	res "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/res_models"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/utils/validation"
	"github.com/gofiber/fiber/v2"
)

type JobsHandler struct {
	job Job.JobServiceClient
}

func NewJobsHandler(job Job.JobServiceClient) *JobsHandler {
	return &JobsHandler{job: job}
}

func (h *JobsHandler) PostJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int64)
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

	res, err := h.job.PostJob(context.Background(), &Job.PostjobReq{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Skills:      req.Skills,
		TimePeriod:  req.TimePeriod,
		Type:        req.Type,
		Budget:      float32(req.Budget),
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)

}

func (h *JobsHandler) SendProposal(c *fiber.Ctx) error {
	var req req.Proposal

	job_id := c.Query("jobID")
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

	res, err := h.job.SendProposal(context.Background(), &Job.ProposalReq{
		Budget:      req.Budget,
		CoverLetter: req.Coverletter,
		UserId:      user_id,
		JobId:       job_id,
		// Attachments: attachments,
	})

	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetMyJobs(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(int64)
	id := strconv.Itoa(int(user_id))
	res, err := h.job.GetMyJobs(context.Background(), &Job.GetMyJobsReq{UserId: id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJobProposals(c *fiber.Ctx) error {
	jobID := c.Params("job_id")
	user_id := c.Locals("User_id").(string)
	res, err := h.job.GetJobProposals(context.Background(), &Job.GJPReq{JobId: jobID, UserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetOffersForJobByClient(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	res, err := h.job.GetOfferByClient(context.Background(), &Job.GFCReq{UserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) EditJob(c *fiber.Ctx) error {
	var req req.PostJob
	user_id := c.Locals("User_id").(int64)
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

	res, err := h.job.PostJob(context.Background(), &Job.PostjobReq{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Skills:      req.Skills,
		TimePeriod:  req.TimePeriod,
		Type:        req.Type,
		Budget:      float32(req.Budget),
		UserId:      user_id,
	})
	if err != nil {
		return c.Status(403).JSON(err.Error())
	}

	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) SendOffer(c *fiber.Ctx) error {
	var req req.SendOffer
	user_id := c.Locals("User_id").(int64)

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
	startDate, err := time.Parse("2-1-2006", req.Starting_time)
	if err != nil {
		c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	Error, err := validation.Validation(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": Error})
	}
	res, err := h.job.SendOffer(context.Background(), &Job.SendOfferReq{
		Budget:       req.Budget,
		OfferLetter:  req.Offer_letter,
		StartingTime: startDate.Format("2-1-2006"),
		JobId:        int32(req.Job_id),
		FreelancerId: int32(req.Freelancer_id),
		ClientId:     int32(user_id),
	})
	if err != nil {
		println(err)
		return c.Status(500).JSON(fiber.Map{"error": "error in rpc connection"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) AcceptOffer(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)
	of_id := c.Params("offer_id")

	res, err := h.job.AcceptOffer(context.Background(), &Job.AcceptOfferReq{UserId: user_id, OfferID: of_id})
	if err != nil {
		print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJobs(c *fiber.Ctx) error {
	res, err := h.job.GetJobs(context.Background(), &Job.NoParam{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) GetJob(c *fiber.Ctx) error {
	job_id := c.Params("id")
	println(job_id)
	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) SendInvoice(c *fiber.Ctx) error {
	var req req.SendInvoice
	user_id := c.Locals("User_id").(string)
	res, err := h.job.SendWeeklyInvoice(context.Background(), &Job.InvoiceReq{ContractID: int32(req.ContractId), TotalHourWorked: float32(req.TotalHoursWorked), SuserId: user_id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(int(res.Status)).JSON(res)
}

func (h *JobsHandler) OnboardFreelancersToPaypal(c *fiber.Ctx) error {
	user_id := c.Locals("User_id").(string)

	payload := createPartnerReferralPayload(user_id)

	// Step 3: Send the HTTP request to PayPal
	accessToken := "A21AAJyH-8bjFx7neRjX8nMBXw9kyFDglypACRr8BhURXiZtwM60nb5hkPEQB4x-CDWnUGJefeGNQUmIiOhjY36iLm1ExA_Eg"
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }
	url := "https://api-m.sandbox.paypal.com/v2/customer/partner-referrals"
	resp, err := sendPayPalRequest(url, accessToken, payload)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Step 4: Handle the response from PayPal
	if resp.StatusCode == http.StatusOK {
		return c.Status(http.StatusOK).JSON(resp.Body)
	} else {
		return c.Status(resp.StatusCode).JSON(resp.Body)
	}
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func generateAccessToken() (string, error) {
	clientID := "AXkOdHFP9XEz40vzsctBXR3p5s5MLwPIpLChTneph9S1yw1e_RwAFECvZhKdPh0-Zak1Jku1tWPbhCfw"
	clientSecret := "EMHI4OpfbpaGrz9-LW0EfhGmbU1C5fcUWocv6Jrt0L-TKkPsS73FSErly_mGd-3NZ1zIuXGila7EDMnw"

	req, err := http.NewRequest("POST", "https://api-m.sandbox.paypal.com/v1/oauth2/token", bytes.NewBufferString("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to generate access token: %s", body)
	}

	// Parse the response body to extract the access token
	parts := strings.Split(string(body), "&")
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if kv[0] == `\"access_token\"` {
			return kv[1], nil
		}
	}

	return "", fmt.Errorf("access token not found in response: %s", body)
}

func sendPayPalRequest(url, accessToken string, payload map[string]interface{}) (*http.Response, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	return client.Do(req)
}

func createPartnerReferralPayload(trackingID string) map[string]interface{} {
	return map[string]interface{}{
		"tracking_id": trackingID,
		"operations": []map[string]interface{}{
			{
				"operation": "API_INTEGRATION",
				"api_integration_preference": map[string]interface{}{
					"rest_api_integration": map[string]interface{}{
						"integration_method": "PAYPAL",
						"integration_type":   "THIRD_PARTY",
						"third_party_details": map[string]interface{}{
							"features": []string{
								"PAYMENT",
								"REFUND",
							},
						},
					},
				},
			},
		},
		"products": []string{
			"EXPRESS_CHECKOUT",
		},
		"legal_consents": []map[string]interface{}{
			{
				"type":    "SHARE_DATA_CONSENT",
				"granted": true,
			},
		},
		"contact": map[string]interface{}{
			"name": map[string]interface{}{
				"given_name": "sellerInfo.Name",
			},
			"email_address": "sellerInfo.Email",
			"website": map[string]interface{}{
				"urls": []string{
					"sellerInfo.BusinessUR",
				},
			},
		},
	}
}

// func (h *JobsHandler) ExecutePaymentForContractWithInvoiceID(c *fiber.Ctx) error {
// 	invoiceID := c.Params("invoice_id")
// 	user_id := c.Locals("user_id")

// 	res, err := h.job.ExecutePayment(context.Background(), &Job.AcceptOfferReq{})
// 	if err != nil {
// 		return c.Status(res.Status).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Render("index.html", res)
// }

// func (h *JobsHandler) CloseJobPost() {

// }

// func (h *JobsHandler) GetContractDetails(c *fiber.Ctx){
// 	job_id:=c.Params("id")
// 	println(job_id)
// 	res, err := h.job.GetJob(context.Background(), &Job.GetJobReq{JobId: job_id})
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
// 	}
// 	return c.Status(int(res.Status)).JSON(res)
// }


func (h *JobsHandler) Search(c *fiber.Ctx) error{
	query:=c.Query("q")
	PayType:=c.Query("t")
	Hourly_rate:=c.Query("hourly_rate")
	fixed_rate:=c.Query("fixed_rate")
	category:=c.Query("c")

	res,err:=h.job.SearchJobs(context.Background(),&Job.SearchJobsReq{Query: query,Paytype: PayType,HourlyRate: Hourly_rate,FixedRate: fixed_rate,Category: category})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(int(res.Status)).JSON(res)

}

//contracts

// func (h *JobsHandler) GetAllContractsForClient(c *fiber.Ctx){
// 	user_id:=c.Locals("User_id").(int64)
// 	res,err:=h.job.GetAllContractsForClient(context.Background(),&J)
// }

