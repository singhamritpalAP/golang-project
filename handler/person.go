package handler

import (
	"github.com/gin-gonic/gin"
	"golang-project/golang-project/service"
	"golang-project/golang-project/utils"
	"log"
	"net/http"
)

type Handler struct {
	Service service.PersonService
}

func (handler *Handler) Get(ctx *gin.Context) {
	// fetch person id from req
	personId, err := utils.GetPersonId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person_id"})
		return
	}
	log.Println("request received for person id: ", personId)
	personDetails, err := handler.Service.GetPerson(personId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, personDetails)
}

func (handler *Handler) Create(context *gin.Context) {

}
