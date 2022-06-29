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
)

type JobOfferHandler struct {
	Service *service.JobOfferService
}

func (handler *JobOfferHandler) AddJobOffer(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "POST /jobOffers")
	defer span.Finish()

	var jobOfferDTO dto.JobOfferRequestDTO
	if err := ctx.ShouldBindJSON(&jobOfferDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	dto, err := handler.Service.Add(&jobOfferDTO)
	if err != nil {
		fmt.Println(err)
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
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	offersDTO, err := handler.Service.GetCompanysOffers(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) GetAll(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers")
	defer span.Finish()

	offersDTO, err := handler.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) Search(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "GET /jobOffers/search")
	defer span.Finish()

	param := ctx.Query("param")
	offersDTO, err := handler.Service.Search(param)
	if err != nil {
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
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	offersDTO, err := handler.Service.GetById(id)
	if err != nil {
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
		ctx.JSON(http.StatusBadRequest, idErr.Error())
		return
	}

	err := handler.Service.Delete(id)
	if err != nil {
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
