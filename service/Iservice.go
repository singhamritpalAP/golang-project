package service

import "golang-project/golang-project/models"

type IService interface {
	GetPerson(personId int) (models.UserData, error)
	CreatePerson()
}
