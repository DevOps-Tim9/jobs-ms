package dto

import "github.com/go-playground/validator"

type JobOfferRequestDTO struct {
	CompanyID                  int    `validate:"required"`
	Position                   string `validate:"required"`
	JobDescription             string `validate:"required"`
	DailyActivitiesDescription string `validate:"required"`
	Skills                     string `validate:"required"`
	Link                       string `validate:"required"`
}

func (u *JobOfferRequestDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
