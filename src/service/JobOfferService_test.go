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

func (suite *JobOfferServiceUnitTestsSuite) TestJobOfficeService_GetById_GetsJobOffer() {
	id := 1
	description := "Description test"
	companyId := 1
	position := "position test"
	skills := "test skills"
	link := "test link"
	activities := "activities test"

	offer := model.JobOffer{
		ID:                         id,
		JobDescription:             description,
		CompanyID:                  companyId,
		Position:                   position,
		Skills:                     skills,
		Link:                       link,
		DailyActivitiesDescription: activities,
	}

	suite.offerRepositoryMock.On("GetById", 1).Return(&offer, nil).Once()

	dto, err := suite.service.GetById(1)

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), description, dto.JobDescription)
	assert.Equal(suite.T(), skills, dto.Skills)
	assert.Equal(suite.T(), position, dto.Position)
	assert.Equal(suite.T(), activities, dto.DailyActivitiesDescription)
	assert.Equal(suite.T(), link, dto.Link)
	assert.Equal(suite.T(), companyId, dto.CompanyID)
}

func (suite *JobOfferServiceUnitTestsSuite) TestJobOfficeService_Delete_Pass() {
	id := 1
	suite.offerRepositoryMock.On("Delete", id).Return(nil).Once()

	err := suite.service.Delete(id)

	assert.Equal(suite.T(), nil, err)
}

func (suite *JobOfferServiceUnitTestsSuite) TestJobOfficeService_GetCompanysOffers_NoOffersReturnsEmpty() {
	suite.offerRepositoryMock.On("GetByCompany", 1).Return([]*model.JobOffer{}, nil).Once()

	offers, err := suite.service.GetCompanysOffers(1)

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 0, len(offers))
}
func (suite *JobOfferServiceUnitTestsSuite) TestJobOfficeService_GetAll_NoOffersReturnsEmpty() {
	suite.offerRepositoryMock.On("GetAll").Return([]*model.JobOffer{}, nil).Once()

	offers, err := suite.service.GetAll()

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), 0, len(offers))
}

func (suite *JobOfferServiceUnitTestsSuite) TestJobOfficeService_Search_ReturnedOffer() {
	offer := model.JobOffer{
		CompanyID:                  1,
		Skills:                     "skills",
		JobDescription:             "test",
		DailyActivitiesDescription: "desc",
		Position:                   "pos",
		Link:                       "link",
		ID:                         1,
	}
	var list []*model.JobOffer
	list = append(list, &offer)
	param := "test"

	suite.offerRepositoryMock.On("Search", param).Return(list, nil).Once()

	offers, err := suite.service.Search(param)

	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), len(list), len(offers))
	for i := 0; i < len(offers); i++ {
		assert.Equal(suite.T(), list[i].ID, offers[i].ID)
		assert.Equal(suite.T(), list[i].CompanyID, offers[i].CompanyID)
		assert.Equal(suite.T(), list[i].Link, offers[i].Link)
		assert.Equal(suite.T(), list[i].Skills, offers[i].Skills)
		assert.Equal(suite.T(), list[i].JobDescription, offers[i].JobDescription)
		assert.Equal(suite.T(), list[i].DailyActivitiesDescription, offers[i].DailyActivitiesDescription)
	}
}
