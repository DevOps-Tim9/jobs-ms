package repository

import (
	"jobs-ms/src/model"

	"github.com/stretchr/testify/mock"
)

type JobOfferRepositoryMock struct {
	mock.Mock
}

func (repo *JobOfferRepositoryMock) Add(offer model.JobOffer) (model.JobOffer, error) {
	args := repo.Called(offer)
	if args.Get(1) == nil {
		return args.Get(0).(model.JobOffer), nil
	}
	return args.Get(0).(model.JobOffer), args.Get(1).(error)
}

func (repo *JobOfferRepositoryMock) GetByCompany(id int) ([]*model.JobOffer, error) {
	args := repo.Called(id)
	if args.Get(1) == nil {
		return args.Get(0).([]*model.JobOffer), nil
	}
	return args.Get(0).([]*model.JobOffer), args.Get(1).(error)
}

func (repo *JobOfferRepositoryMock) GetAll() ([]*model.JobOffer, error) {
	args := repo.Called()
	if args.Get(1) == nil {
		return args.Get(0).([]*model.JobOffer), nil
	}
	return args.Get(0).([]*model.JobOffer), args.Get(1).(error)
}

func (repo *JobOfferRepositoryMock) Search(param string) ([]*model.JobOffer, error) {
	args := repo.Called(param)
	if args.Get(1) == nil {
		return args.Get(0).([]*model.JobOffer), nil
	}
	return args.Get(0).([]*model.JobOffer), args.Get(1).(error)
}
