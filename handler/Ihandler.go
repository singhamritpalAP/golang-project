package handler

import "github.com/gin-gonic/gin"

type IHandler interface {
	Get(ctx *gin.Context)
	Create(context *gin.Context)
}
