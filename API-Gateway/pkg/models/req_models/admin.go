package req

type AddSkills struct{
	Skill string `json:"skill" validate:"required"`
	Description string `json:"description" validate:"required"`
}