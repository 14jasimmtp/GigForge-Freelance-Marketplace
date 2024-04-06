package req

type AddSkills struct{
	Skill string `json:"skill"`
	Description string `json:"description"`
	Category int64 `json:"category"`
	
}