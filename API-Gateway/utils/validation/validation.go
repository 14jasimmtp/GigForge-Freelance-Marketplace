package validation

import (
	"errors"
	"fmt"

	req "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/pkg/models/req_models"
	"github.com/go-playground/validator"
)

func Validation(data interface{}) (*[]req.Errors, error) {
	var afterErrorCorection []req.Errors
	var result req.Errors
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Tag() {
				case "required":
					err := fmt.Sprintf("%s is required", e.Field())
					result = req.Errors{Error: err}
				case "min":
					err := fmt.Sprintf("%s should be at least %s", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "max":
					err := fmt.Sprintf("%s should be at most %s", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "email":
					err := fmt.Sprintf("%s should be email structure %s ", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "eqfield":
					err := fmt.Sprintf("%s should be equal with %s ", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "len":
					err := fmt.Sprintf("%s should be have  %s ", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "alpha":
					err := fmt.Sprintf("%s should be Alphabet ", e.Field())
					result = req.Errors{Error: err}
				case "number":
					err := fmt.Sprintf("%s should be numeric %s ", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "numeric":
					err := fmt.Sprintf("%s should be  numeric %s ", e.Field(), e.Param())
					result = req.Errors{Error: err}
				case "uppercase":
					err := fmt.Sprintf("%s should be  %s %s ", e.Field(), e.Tag(), e.Param())
					result = req.Errors{Error: err}
				case "regexp":
					err := fmt.Sprintf("%s should contain atleast one %s",e.Field(),e.Param())
					result = req.Errors{Error: err}
				}

				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, errors.New("doesn't fulfill the requirements")
	}
	return &afterErrorCorection, nil
}
