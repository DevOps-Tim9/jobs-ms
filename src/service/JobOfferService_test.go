package service

import (
	"jobs-ms/src/dto"
	"jobs-ms/src/model"
	"jobs-ms/src/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JobOfferServiceUnitTestsSuite struct {
	suite.Suite
	offerRepositoryMock *repository.JobOfferRepositoryMock
	service             IJobOfferService
}

func TestJobOfferServiceUnitTestsSuite(t *testing.T) {
	suite.Run(t, new(JobOfferServiceUnitTestsSuite))
}

func (suite *JobOfferServiceUnitTestsSuite) SetupSuite() {
	suite.offerRepositoryMock = new(repository.JobOfferRepositoryMock)
	suite.service = NewJobOfferService(suite.offerRepositoryMock)
}

func (suite *JobOfferServiceUnitTestsSuite) TestNewJobOfferService() {
	assert.NotNil(suite.T(), suite.service, "Service is nil")
}

func (suite *JobOfferServiceUnitTestsSuite) TestJobOfferService_Add_ValidDataProvided() {
	dto := dto.JobOfferRequestDTO{
		CompanyID:                  1,
		Skills:                     "skills",
		JobDescription:             "desc",
		DailyActivitiesDescription: "desc",
		Position:                   "pos",
		Link:                       "link",
	}

	entity := model.JobOffer{
		CompanyID:                  1,
		Skills:                     "skills",
		JobDescription:             "desc",
		DailyActivitiesDescription: "desc",
		Position:                   "pos",
		Link:                       "link",
	}

	savedEntity := model.JobOffer{
		CompanyID:                  1,
		Skills:                     "skills",
		JobDescription:             "desc",
		DailyActivitiesDescription: "desc",
		Position:                   "pos",
		Link:                       "link",
		ID:                         1,
	}

	suite.offerRepositoryMock.On("Add", entity).Return(savedEntity, nil).Once()

	returnedOffer, err := suite.service.Add(&dto)

	assert.Equal(suite.T(), dto.CompanyID, returnedOffer.CompanyID)
	assert.Equal(suite.T(), dto.Link, returnedOffer.Link)
	assert.Equal(suite.T(), dto.Position, returnedOffer.Position)
	assert.Equal(suite.T(), dto.JobDescription, returnedOffer.JobDescription)
	assert.Equal(suite.T(), dto.DailyActivitiesDescription, returnedOffer.DailyActivitiesDescription)
	assert.Equal(suite.T(), dto.Skills, returnedOffer.Skills)
	assert.Equal(suite.T(), nil, err)
}
