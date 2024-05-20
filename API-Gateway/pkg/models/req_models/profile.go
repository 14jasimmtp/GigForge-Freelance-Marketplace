package req


type Education struct {
	School        string `json:"school" validate:"required"`
	Course        string `json:"Course" validate:"required"`
	Date_Started  string `json:"date_started" validate:"required"`
	Date_Ended    string `json:"date_ended" validate:"required"`
	Area_Of_Study string `json:"area_of_study" validate:"required"`
	Description   string `json:"description" validate:"required"`
}

type Profile struct {
	Title       string `json:"Title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Hourly_rate int64  `json:"hourly_rate" validate:"required"`
}

type Skills struct {
	Skills []int64 `json:"skills" validate:"required"`
}

type Experience struct {
	Company  string    `json:"company" validate:"required"`
	City     string    `json:"city" validate:"required"`
	Country  string    `json:"country" validate:"required"`
	Title    string    `json:"title" validate:"required"`
	FromDate string    `json:"from-date" validate:"required"`
	ToDate   string    `json:"to-date" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type AddReview struct{
	Review string `json:"review" validate:"required"`
	Rating int `json:"rating" validate:"required,numeric,gte=1,lte=5"`
	Freelancer_id int `json:"freelancer_id" validate:"required"`
}



