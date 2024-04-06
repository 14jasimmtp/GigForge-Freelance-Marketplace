package req

import (
	"mime/multipart"
)

type PostJob struct {
	Title       string  `json:"title"`
	Description string  `json:"Description"`
	Skills      []int64 `json:"skills"`
	TimePeriod  string  `json:"time-period"`
	Level       string  `json:"freelancer_level"`
	Category    int64   `json:"category"`
	Budget      string  `json:"budget"`
}

type Proposal struct {
	Budget      string `json:"budget"`
	Coverletter string `json:"cover_letter"`
	Attachments []*multipart.FileHeader `form:"attachments"`
}
