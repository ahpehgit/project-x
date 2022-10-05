package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func NewMainController() *MainController {
	return &MainController{}
}

func (controller *MainController) Root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
