package req

import (
	"mime/multipart"
	"time"
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
	Budget        float32
	Offer_letter  string
	Starting_time time.Time
	Job_id int
	Freelancer_id int
}
