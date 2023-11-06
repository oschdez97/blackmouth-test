package app

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/oschdez97/blackmouth-test/controller"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/oschdez97/blackmouth-test/docs"
)

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (a *App) CreateConnection(){
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}


func (a *App) CreateRoutes() {
	routes := gin.Default()
	
	routes.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	api := routes.Group("/api/v1")
	
	playerController := controller.NewPlayerController(a.DB)
	api.GET("/player", playerController.GetPlayer)
	api.GET("/player/:id", playerController.GetPlayerById)
	api.POST("/player", playerController.InsertPlayer)
	api.PUT("/player/:id", playerController.UpdatePlayer)
	api.DELETE("/player/:id", playerController.DeletePlayer)

	queueController := controller.NewQueueController(a.DB, QUEUE_COUNT)
	api.GET("/queue", queueController.GetQueue)
	api.GET("/queue/:id", queueController.GetQueueById)
	api.POST("/queue", queueController.InsertQueue)
	api.PUT("/queue/:id", queueController.UpdateQueue)
	api.DELETE("/queue/:id", queueController.DeleteQueue)

	gameSessionController := controller.NewGameSessionController(a.DB)
	api.GET("/gamesession", gameSessionController.GetGameSession)
	api.GET("/gamesession/:status", gameSessionController.GetGameSessionByStatus)
	api.POST("/gamesession", gameSessionController.JoinGameSession)

	a.Routes = routes
}

func (a *App) Run(){
	a.Routes.Run(":8080")
}
