package service

import (
	"golang-project/golang-project/models"
	"golang-project/golang-project/relationaldatabase"
	"log"
)

type PersonService struct {
}

func (service *PersonService) GetPerson(personId int) (models.UserData, error) {
	// Query the database for the person
	person, err := relationaldatabase.Get(personId)
	if err != nil {
		log.Println("Error getting person from database")
		return models.UserData{}, err
	}
	return person, nil

}

func (service *PersonService) CreatePerson() {

}
