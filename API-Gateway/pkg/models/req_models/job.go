package req

import (
	"mime/multipart"
)

type PostJob struct {
	Title       string  `json:"title"`
	Description string  `json:"Description"`
	Skills      []int64 `json:"skills"`
	TimePeriod  string  `json:"time-period"`
	Type        string  `json:"type"`
	Category    int64   `json:"category"`
	Budget      float64 `json:"budget"`
}

type Proposal struct {
	Budget      float32                 `json:"budget" validate:"required"`
	Coverletter string                  `json:"cover_letter" validate:"required"`
	Attachments []*multipart.FileHeader `form:"attachments" `
}

type SendOffer struct {
	Budget        float32 `json:"budget"`
	Offer_letter  string  `json:"offer_letter"`
	Starting_time string `json:"starting_time"`
	Job_id        int `json:"job_id"`
	Freelancer_id int `json:"freelancer_id"`
}

type SendInvoice struct{
	ContractId int `json:"contractID"`
	TotalHoursWorked int `json:"Hours_worked"`

}
