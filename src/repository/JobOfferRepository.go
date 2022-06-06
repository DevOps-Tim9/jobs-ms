package repository

import (
	"errors"
	"jobs-ms/src/model"
	"strings"

	"github.com/jinzhu/gorm"
)

type IJobOfferRepository interface {
	Add(model.JobOffer) (model.JobOffer, error)
	GetByCompany(int) ([]*model.JobOffer, error)
	GetAll() ([]*model.JobOffer, error)
	Search(string) ([]*model.JobOffer, error)
}

func NewJobOfferRepository(database *gorm.DB) IJobOfferRepository {
	return &JobOfferRepository{
		database,
	}
}

type JobOfferRepository struct {
	Database *gorm.DB
}

func (repo *JobOfferRepository) Add(offer model.JobOffer) (model.JobOffer, error) {
	err := repo.Database.Save(&offer).Error

	return offer, err
}

func (repo *JobOfferRepository) GetByCompany(id int) ([]*model.JobOffer, error) {
	var offers = []*model.JobOffer{}
	if result := repo.Database.Find(&offers, "company_id = ?", id); result.Error != nil {
		return nil, errors.New("Error happened during retrieving company's job offers")
	}

	return offers, nil
}

func (repo *JobOfferRepository) GetAll() ([]*model.JobOffer, error) {
	var offers = []*model.JobOffer{}
	if result := repo.Database.Find(&offers); result.Error != nil {
		return nil, errors.New("Error happened during retrieving company's job offers")
	}

	return offers, nil
}

func (repo *JobOfferRepository) Search(param string) ([]*model.JobOffer, error) {
	searchParam := "%" + strings.ToLower(param) + "%"
	var offers = []*model.JobOffer{}
	if result := repo.Database.Find(&offers, "LOWER(position) LIKE $1 OR LOWER(skills) LIKE $1 OR LOWER(daily_activities_description) LIKE $1 OR LOWER(job_description) LIKE $1", searchParam); result.Error != nil {
		return nil, errors.New("Error happened during retrieving company's job offers")
	}

	return offers, nil
}
