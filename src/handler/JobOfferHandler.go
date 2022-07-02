package handler

import (
	"errors"
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type JobOfferHandler struct {
	Service *service.JobOfferService
	Logger  *logrus.Entry
}

func (handler *JobOfferHandler) AddJobOffer(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "POST /jobOffers")
	defer span.Finish()

	var jobOfferDTO dto.JobOfferRequestDTO
	if err := ctx.ShouldBindJSON(&jobOfferDTO); err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	handler.Logger.Info(fmt.Sprintf("Adding new job offer for company %d", jobOfferDTO.CompanyID))

	dto, err := handler.Service.Add(&jobOfferDTO)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto)
}

func (handler *JobOfferHandler) GetJobOffersByCompany(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers/company/:companyId")
	defer span.Finish()

	id, idErr := getId(ctx.Param("companyId"))
	if idErr != nil {
		handler.Logger.Debug(idErr.Error())
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}
	handler.Logger.Info(fmt.Sprintf("Getting job offers for company %d", id))
	offersDTO, err := handler.Service.GetCompanysOffers(id)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) GetAll(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers")
	defer span.Finish()

	handler.Logger.Info("Getting job offers")

	offersDTO, err := handler.Service.GetAll()
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) Search(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers/search")
	defer span.Finish()

	handler.Logger.Info("Searching job offers")

	param := ctx.Query("param")
	offersDTO, err := handler.Service.Search(param)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) GetJobOffer(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers/:id")
	defer span.Finish()

	id, idErr := getId(ctx.Param("id"))
	if idErr != nil {
		handler.Logger.Debug(idErr.Error())
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	handler.Logger.Info(fmt.Sprintf("Getting job offer with id %d", id))

	offersDTO, err := handler.Service.GetById(id)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) DeleteJobOffer(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "DELETE /jobOffers/:id")
	defer span.Finish()

	id, idErr := getId(ctx.Param("id"))
	if idErr != nil {
		handler.Logger.Debug(idErr.Error())
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	handler.Logger.Info(fmt.Sprintf("Deleting job offer with id %d", id))

	err := handler.Service.Delete(id)
	if err != nil {
		handler.Logger.Debug(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func getId(idParam string) (int, error) {
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return 0, errors.New("Company id should be a number")
	}
	return int(id), nil
}
