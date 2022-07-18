package mapper

import (
	"jobs-ms/src/dto"
	"jobs-ms/src/model"
)

func JobOfferToJobOfferResponseDTO(jobOffer *model.JobOffer) *dto.JobOfferResponseDTO {
	var offer dto.JobOfferResponseDTO

	offer.ID = jobOffer.ID
	offer.CompanyID = jobOffer.CompanyID
	offer.JobDescription = jobOffer.JobDescription
	offer.DailyActivitiesDescription = jobOffer.DailyActivitiesDescription
	offer.Link = jobOffer.Link
	offer.Position = jobOffer.Position
	offer.Skills = jobOffer.Skills

	return &offer
}

func JobOfferRequestDTOToJobOffer(jobOffer *dto.JobOfferRequestDTO) *model.JobOffer {
	var offer model.JobOffer

	offer.CompanyID = jobOffer.CompanyID
	offer.JobDescription = jobOffer.JobDescription
	offer.DailyActivitiesDescription = jobOffer.DailyActivitiesDescription
	offer.Link = jobOffer.Link
	offer.Position = jobOffer.Position
	offer.Skills = jobOffer.Skills

	return &offer
}
