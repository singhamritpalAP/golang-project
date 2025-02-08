package service

import (
	"golang-project/golang-project/models"
	"golang-project/golang-project/relationaldatabase"
	"log"
)

type PersonService struct {
}

// GetPerson for fetching person details based on personId
func (service *PersonService) GetPerson(personId int) (models.UserData, error) {
	log.Println("received request for person with id ", personId)
	// Query the database for the person
	person, err := relationaldatabase.Get(personId)
	if err != nil {
		log.Println("Error getting person from database")
		return models.UserData{}, err
	}
	return person, nil

}

// CreatePerson for storing person details in database
func (service *PersonService) CreatePerson(userData models.UserData) error {
	log.Println("Received request to create person")
	return relationaldatabase.CreateUser(userData)
}
