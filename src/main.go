package main

import (
	"fmt"
	"io"
	"jobs-ms/src/handler"
	"jobs-ms/src/model"
	"jobs-ms/src/repository"
	"jobs-ms/src/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
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
	db, _ = gorm.Open("postgres", connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(model.JobOffer{})
	return db, err
}

func InitJaeger() (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: "jobs-ms",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	return tracer, closer, err
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
	router.GET("/jobOffers/company/:companyId", handler.GetJobOffersByCompany)
	router.GET("/jobOffers/search", handler.Search)
	router.GET("/jobOffers/:id", handler.GetJobOffer)
	router.DELETE("/jobOffers/:id", handler.DeleteJobOffer)
}

func main() {
	database, _ := initDB()

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	tracer, trCloser, err := InitJaeger()
	if err != nil {
		fmt.Printf("error init jaeger %v", err)
	} else {
		defer trCloser.Close()
		opentracing.SetGlobalTracer(tracer)
	}

	offerRepo := initOfferRepo(database)
	offerService := initOfferService(offerRepo)
	offerHandler := initOfferHandler(offerService)

	router := gin.Default()

	handleOfferFunc(offerHandler, router)

	// http.ListenAndServe(port, cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:9094"},
	// 	AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	// }).Handler(router))
	http.ListenAndServe(port, cors.AllowAll().Handler(router))
}
