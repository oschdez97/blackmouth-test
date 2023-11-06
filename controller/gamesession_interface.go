package controller

import "github.com/gin-gonic/gin"

type GameSessionControllerInterface interface {
	GetGameSession(g *gin.Context)
	GetGameSessionByStatus(g *gin.Context)
	JoinGameSession(g *gin.Context)
}