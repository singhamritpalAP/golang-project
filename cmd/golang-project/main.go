package main

import (
	"github.com/gin-gonic/gin"
	"golang-project/golang-project/handler"
	"golang-project/golang-project/service"
	"log"
)

func main() {
	router := gin.Default()

	// Initialize the service todo initialise service and handlers in respective files also initialize db
	personService := service.PersonService{}

	// Initialize the handler with the service
	handlers := &handler.Handler{Service: personService}

	// GET /person/:person_id/info
	router.GET("/person/:person_id/info", handlers.Get)

	// POST /person/create
	router.POST("/person/create", handlers.Create)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
