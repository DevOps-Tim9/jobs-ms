package service

import (
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/mapper"
	"jobs-ms/src/repository"
)

type JobOfferService struct {
	JobOfferRepo repository.IJobOfferRepository
}

type IJobOfferService interface {
	Add(*dto.JobOfferRequestDTO) (*dto.JobOfferResponseDTO, error)
}

func NewJobOfferService(jobOfferRepository repository.IJobOfferRepository) IJobOfferService {
	return &JobOfferService{
		jobOfferRepository,
	}
}

func (service *JobOfferService) Add(dto *dto.JobOfferRequestDTO) (*dto.JobOfferResponseDTO, error) {
	err := dto.Validate()
	if err != nil {
		return nil, err
	}

	entity := mapper.JobOfferRequestDTOToJobOffer(dto)

	addedEntity, err := service.JobOfferRepo.Add(*entity)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return mapper.JobOfferToJobOfferResponseDTO(&addedEntity), nil
}
