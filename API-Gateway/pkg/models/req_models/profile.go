package req

type Education struct {
	School        string `json:"school"`
	Course        string `json:"Course"`
	Date_Started  string `json:"date_started"`
	Date_Ended    string `json:"date_ended"`
	Area_Of_Study string `json:"area_of_study"`
	Description   string `json:"description"`
}

type Profile struct{
	Title string `json:"Title"`
	Description string `json:"description"`
	Hourly_rate int64 `json:"hourly_rate"`
}