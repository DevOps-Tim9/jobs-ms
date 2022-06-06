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
	GetCompanysOffers(int) ([]*dto.JobOfferResponseDTO, error)
	GetAll() ([]*dto.JobOfferResponseDTO, error)
	Search(string) ([]*dto.JobOfferResponseDTO, error)
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

func (service *JobOfferService) GetCompanysOffers(id int) ([]*dto.JobOfferResponseDTO, error) {
	offers, err := service.JobOfferRepo.GetByCompany(id)

	if err != nil {
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	return res, nil
}

func (service *JobOfferService) GetAll() ([]*dto.JobOfferResponseDTO, error) {
	offers, err := service.JobOfferRepo.GetAll()

	if err != nil {
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	return res, nil
}

func (service *JobOfferService) Search(param string) ([]*dto.JobOfferResponseDTO, error) {
	offers, err := service.JobOfferRepo.Search(param)

	if err != nil {
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	return res, nil
}
