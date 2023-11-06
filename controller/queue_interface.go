package controller

import "github.com/gin-gonic/gin"

type QueueControllerInterface interface {
	InsertQueue(g *gin.Context)
	GetQueue(g *gin.Context)
	GetQueueById(g *gin.Context)
	UpdateQueue(g *gin.Context)
	DeleteQueue(g *gin.Context)
}