package handler

import (
	"errors"
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobOfferHandler struct {
	Service *service.JobOfferService
}

func (handler *JobOfferHandler) AddJobOffer(ctx *gin.Context) {
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
	offersDTO, err := handler.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func (handler *JobOfferHandler) Search(ctx *gin.Context) {
	param := ctx.Query("param")
	offersDTO, err := handler.Service.Search(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, offersDTO)
}

func getId(idParam string) (int, error) {
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return 0, errors.New("Company id should be a number")
	}
	return int(id), nil
}
