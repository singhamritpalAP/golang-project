package utils

import (
	"github.com/gin-gonic/gin"
	"golang-project/golang-project/constants"
	"strconv"
)

// function to fetch person id from request
func GetPersonId(ctx *gin.Context) (int, error) {
	personIdStr := ctx.Param(constants.PersonKey)
	personId, err := strconv.Atoi(personIdStr)
	if err != nil {
		return 0, err
	}
	return personId, nil
}
