package main

import (
	"fmt"
	"jobs-ms/src/handler"
	"jobs-ms/src/model"
	"jobs-ms/src/repository"
	"jobs-ms/src/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

func initDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB")

	connString := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s sslmode=disable password=%s", user, dbName, pass)
	db, err = gorm.Open("postgres", connString)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.JobOffer{})

	return db
}

func initOfferRepo(database *gorm.DB) *repository.JobOfferRepository {
	return &repository.JobOfferRepository{Database: database}
}

func initOfferService(repo *repository.JobOfferRepository) *service.JobOfferService {
	return &service.JobOfferService{JobOfferRepo: repo}
}

func initOfferHandler(service *service.JobOfferService) *handler.JobOfferHandler {
	return &handler.JobOfferHandler{Service: service}
}

func handleOfferFunc(handler *handler.JobOfferHandler, router *gin.Engine) {
	router.POST("/jobOffers", handler.AddJobOffer)
	router.GET("/jobOffers", handler.GetAll)
	router.GET("/jobOffers/:companyId", handler.GetJobOffersByCompany)
	router.GET("/jobOffers/search", handler.Search)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := initDB()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	offerRepo := initOfferRepo(database)
	offerService := initOfferService(offerRepo)
	offerHandler := initOfferHandler(offerService)

	router := gin.Default()

	handleOfferFunc(offerHandler, router)

	http.ListenAndServe(port, cors.AllowAll().Handler(router))
}
