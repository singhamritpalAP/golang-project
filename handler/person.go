package handler

import (
	"github.com/gin-gonic/gin"
	"golang-project/golang-project/models"
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

func (handler *Handler) Create(ctx *gin.Context) {
	var userData models.UserData

	if err := ctx.BindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("request received for user: ", userData)
	err := handler.Service.CreatePerson(userData)
	if err != nil {
		log.Println("error creating person: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.Status(http.StatusOK)
}
