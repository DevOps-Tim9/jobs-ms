package service

import (
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/mapper"
	"jobs-ms/src/repository"

	"github.com/sirupsen/logrus"
)

type JobOfferService struct {
	JobOfferRepo repository.IJobOfferRepository
	Logger       *logrus.Entry
}

type IJobOfferService interface {
	Add(*dto.JobOfferRequestDTO) (*dto.JobOfferResponseDTO, error)
	GetCompanysOffers(int) ([]*dto.JobOfferResponseDTO, error)
	GetAll() ([]*dto.JobOfferResponseDTO, error)
	Search(string) ([]*dto.JobOfferResponseDTO, error)
	GetById(int) (*dto.JobOfferResponseDTO, error)
	Delete(int) error
}

func NewJobOfferService(jobOfferRepository repository.IJobOfferRepository, logger *logrus.Entry) IJobOfferService {
	return &JobOfferService{
		jobOfferRepository,
		logger,
	}
}

func (service *JobOfferService) Add(dto *dto.JobOfferRequestDTO) (*dto.JobOfferResponseDTO, error) {
	err := dto.Validate()
	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	entity := mapper.JobOfferRequestDTOToJobOffer(dto)

	service.Logger.Info("Adding new job offer in database")

	addedEntity, err := service.JobOfferRepo.Add(*entity)
	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	service.Logger.Info(fmt.Sprintf("Successfully added new job offer in database with id %d", addedEntity.ID))
	return mapper.JobOfferToJobOfferResponseDTO(&addedEntity), nil
}

func (service *JobOfferService) GetCompanysOffers(id int) ([]*dto.JobOfferResponseDTO, error) {
	service.Logger.Info(fmt.Sprintf("Getting job offers from database for company %d", id))
	offers, err := service.JobOfferRepo.GetByCompany(id)

	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	service.Logger.Info(fmt.Sprintf("Successfully got job offers from database for company %d", id))
	return res, nil
}

func (service *JobOfferService) GetAll() ([]*dto.JobOfferResponseDTO, error) {
	service.Logger.Info("Getting job offers from database for company ")
	offers, err := service.JobOfferRepo.GetAll()

	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	service.Logger.Info("Successfully got job offers from database")
	return res, nil
}

func (service *JobOfferService) Search(param string) ([]*dto.JobOfferResponseDTO, error) {
	service.Logger.Info(fmt.Sprintf("Searching job offers from database for param %s", param))
	offers, err := service.JobOfferRepo.Search(param)

	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	res := make([]*dto.JobOfferResponseDTO, len(offers))
	for i := 0; i < len(offers); i++ {
		res[i] = mapper.JobOfferToJobOfferResponseDTO(offers[i])
	}

	service.Logger.Info(fmt.Sprintf("Successfully searched job offers from database by param %s", param))
	return res, nil
}

func (service *JobOfferService) GetById(id int) (*dto.JobOfferResponseDTO, error) {
	service.Logger.Info(fmt.Sprintf("Getting job offer from database with id %d", id))
	offer, err := service.JobOfferRepo.GetById(id)

	if err != nil {
		service.Logger.Debug(err.Error())
		return nil, err
	}

	dto := mapper.JobOfferToJobOfferResponseDTO(offer)
	service.Logger.Info(fmt.Sprintf("Successfully got job offer from database with id %d", id))
	return dto, nil
}

func (service *JobOfferService) Delete(id int) error {
	service.Logger.Info(fmt.Sprintf("Deleting job offer from database with id %d", id))
	err := service.JobOfferRepo.Delete(id)

	if err != nil {
		service.Logger.Debug(err.Error())
		return err
	}

	service.Logger.Info(fmt.Sprintf("Successfully deleted job offer from database with id %d", id))
	return nil
}
