package repository

import (
	"jobs-ms/src/model"

	"github.com/jinzhu/gorm"
)

type IJobOfferRepository interface {
	Add(model.JobOffer) (model.JobOffer, error)
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
