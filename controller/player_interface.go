package controller

import "github.com/gin-gonic/gin"

type PlayerControllerInterface interface {
	InsertPlayer(g *gin.Context)
	GetPlayer(g *gin.Context)
	GetPlayerById(g *gin.Context)
	UpdatePlayer(g *gin.Context)
	DeletePlayer(g *gin.Context)
}