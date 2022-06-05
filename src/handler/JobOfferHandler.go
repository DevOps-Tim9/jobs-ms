package handler

import (
	"fmt"
	"jobs-ms/src/dto"
	"jobs-ms/src/service"
	"net/http"

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
