package req

type AddProject struct{
	Title string `json:"title,omitempty" validate:"required"`
	Description string `json:"Description,omitempty" validate:"required"`
	Category int `json:"category,omitempty" validate:"required"`
	Type string `json:"type,omitempty" validate:"required"`
	Price float64 `json:"price,omitempty"`
	DeliveryDays int64 `json:"Delivery days,omitempty"`
	NumberOfRevisions int64 `json:"Number_of_revisions,omitempty"`
	Starter CustomTier `json:"starter,omitempty"`
	Standard CustomTier `json:"standard,omitempty"`
	Advanced CustomTier `json:"advanced,omitempty"`
}

type CustomTier struct{
	Title string `json:"title,omitempty"`
	Description string `json:"Description,omitempty"`
	Price float64 `json:"project_price"`
	DeliveryDays int64 `json:"Delivery days,omitempty"`
	NumberOfRevisions int64 `json:"Number_of_revisions,omitempty"`
}

