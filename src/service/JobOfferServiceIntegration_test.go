package service

import (
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/model"
	"jobs-ms/src/repository"
	"jobs-ms/src/utils"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JobOfferServiceIntegrationTestSuite struct {
	suite.Suite
	service JobOfferService
	db      *gorm.DB
	offers  []model.JobOffer
}

func (suite *JobOfferServiceIntegrationTestSuite) SetupSuite() {
	host := os.Getenv("DATABASE_DOMAIN")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_SCHEMA")
	port := os.Getenv("DATABASE_PORT")

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		name,
		port,
	)
	db, _ := gorm.Open("postgres", connectionString)

	db.AutoMigrate(model.JobOffer{})

	jobOfferRepository := repository.JobOfferRepository{Database: db}

	suite.db = db

	suite.service = JobOfferService{
		JobOfferRepo: &jobOfferRepository,
		Logger:       utils.Logger(),
	}

	suite.offers = []model.JobOffer{
		{
			CompanyID:                  1000,
			Position:                   "QA",
			JobDescription:             "test",
			DailyActivitiesDescription: "test test test",
			Skills:                     "test",
			Link:                       "test link",
		},
		{
			CompanyID:                  2000,
			Position:                   "QA",
			JobDescription:             "test",
			DailyActivitiesDescription: "test test test",
			Skills:                     "test",
			Link:                       "test link",
		},
	}

	tx := suite.db.Begin()

	tx.Create(&suite.offers[0])

	tx.Commit()

	tx = suite.db.Begin()

	tx.Create(&suite.offers[1])

	tx.Commit()
}

func TestJobOfferServiceIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(JobOfferServiceIntegrationTestSuite))
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_GetById_JobOfferDoesNotExist() {
	id := 2000000

	offer, err := suite.service.GetById(id)

	assert.Nil(suite.T(), offer)
	assert.NotNil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_GetById_JobOfferExists() {
	id := 1

	offer, err := suite.service.GetById(id)

	assert.NotNil(suite.T(), offer)
	assert.Equal(suite.T(), id, offer.ID)
	assert.Nil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_GetAll_JobOffersExist() {
	offers, err := suite.service.GetAll()

	assert.NotNil(suite.T(), offers)
	assert.Equal(suite.T(), 2, len(offers))
	assert.Nil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_GetCompanysOffers_OneJobOfferExists() {
	companyId := 2000

	offers, err := suite.service.GetCompanysOffers(companyId)

	assert.NotNil(suite.T(), offers)
	assert.Equal(suite.T(), 1, len(offers))
	assert.Equal(suite.T(), companyId, offers[0].CompanyID)
	assert.Nil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_GetCompanysOffers_NoJobOffers() {
	companyId := 100000

	offers, err := suite.service.GetCompanysOffers(companyId)

	assert.NotNil(suite.T(), offers)
	assert.Equal(suite.T(), 0, len(offers))
	assert.Nil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_Delete_JobOfferDoesNotExist() {
	id := 2000000

	suite.service.Delete(id)

	assert.True(suite.T(), true)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_Search_JobOfferDoesNotExist() {
	param := "nonexisting param"

	offers, err := suite.service.Search(param)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), offers)
	assert.Equal(suite.T(), 0, len(offers))
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_Search_JobOffersExist() {
	param := "test"

	offers, err := suite.service.Search(param)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), offers)
	assert.Equal(suite.T(), 2, len(offers))
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_Delete_JobOfferExists() {
	id := 2

	err := suite.service.Delete(id)

	assert.Nil(suite.T(), err)
}

func (suite *JobOfferServiceIntegrationTestSuite) TestIntegrationJobOfferService_Add_Pass() {
	offerDto := dto.JobOfferRequestDTO{
		CompanyID:                  1000,
		JobDescription:             "test",
		DailyActivitiesDescription: "test",
		Position:                   "test",
		Skills:                     "test",
		Link:                       "test",
	}

	responseDto, err := suite.service.Add(&offerDto)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), responseDto)
	assert.Equal(suite.T(), offerDto.JobDescription, responseDto.JobDescription)
	assert.Equal(suite.T(), offerDto.DailyActivitiesDescription, responseDto.DailyActivitiesDescription)
	assert.Equal(suite.T(), offerDto.Skills, responseDto.Skills)
	assert.Equal(suite.T(), offerDto.CompanyID, responseDto.CompanyID)
	assert.Equal(suite.T(), offerDto.Position, responseDto.Position)

	suite.service.Delete(responseDto.ID)
}
