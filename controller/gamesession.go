package controller

import (
	"strconv"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/oschdez97/blackmouth-test/model"
	"github.com/oschdez97/blackmouth-test/repository"
)

type GameSessionController struct {
	DB *sql.DB
}

func NewGameSessionController(db *sql.DB) GameSessionControllerInterface {
	return &GameSessionController{DB: db}
}

// GetGameSession    godoc
// @Summary      	 Get gamesession array
// @Description  	 Responds with the list of all gamesessions as JSON.
// @Tags         	 gamesessions
// @Produce      	 json
// @Success      	 200  {array}  model.Player
// @Router       	 /gamesession [get]
func (m *GameSessionController) GetGameSession(g *gin.Context) {
	db := m.DB
	repo_gamesession := repository.NewGameSessionRepository(db)

	get_gamesession := repo_gamesession.SelectGameSession()
	if get_gamesession != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_gamesession, "msg": "get gamesession successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get gamesession successfully"})
	}
}

// GetGameSessionByStatus   godoc
// @Summary      			Get gamesession by status
// @Description  			Responds with the list of gamesessions with the given status as JSON.
// @Tags         			gamesessions
// @Produce      			json
// @Param        			status  path	string  true	"GameSession status"
// @Success      			200  {array}  model.Player
// @Router       			/gamesession/{status} [get]
func (m *GameSessionController) GetGameSessionByStatus(g *gin.Context) {
	db := m.DB
	status64, _ := strconv.ParseUint(g.Param("status"), 10, 32)
	
	repo_gamesession := repository.NewGameSessionRepository(db)

	get_gamesession := repo_gamesession.SelectGameSessionByStatus(uint(status64))
	if get_gamesession != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_gamesession, "msg": "get gamesession successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "gamesession not found"})
	}
}

// JoinGamesession  godoc
// @Summary      	Take a player and join him in a game session
// @Description  	Takes a player ID and a queue ID JSON and determines if there is an open session for that queue, in which case it joins it, otherwise a new session is created.
// @Tags         	gamesessions
// @Produce      	json
// @Param        	joindata  body	model.PostGameSession  true  "PostGameSession JSON"
// @Success      	200  {array}  model.GameSession
// @Router       	/gamesession [post]
func (m *GameSessionController) JoinGameSession(g *gin.Context) {
	db := m.DB
	var post model.PostGameSession
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_player := repository.NewPlayerRepository(db)
		repo_queue := repository.NewQueueRepository(db)
		repo_gamesession := repository.NewGameSessionRepository(db)

		get_player := repo_player.SelectPlayerById(post.PlayerId)
		get_queue := repo_queue.SelectQueueById(post.QueueId)
		
		if get_player != nil && get_queue != nil {
			get_gamesession := repo_gamesession.FindExistingGameSession(post)
			if get_gamesession != nil {
				// update existing gamesession
				updated := repo_gamesession.UpdateGameSession(int(get_gamesession[0].Id), int(get_queue[0].Capacity), post)
				if updated {
					gamesession := repo_gamesession.SelectGameSessionById(uint(get_gamesession[0].Id))
					g.JSON(200, gin.H{"status": "success", "data": gamesession, "msg": "get player successfully"})
					return
				}
			} else {
				// create new gamesession
				new_gamesession := repo_gamesession.InsertGameSession(post)
				if new_gamesession > 0 {
					updated := repo_gamesession.UpdateGameSession(new_gamesession, int(get_queue[0].Capacity), post)
					if updated {
						gamesession := repo_gamesession.SelectGameSessionById(uint(new_gamesession))
						g.JSON(200, gin.H{"status": "success", "data": gamesession, "msg": "get player successfully"})
						return
					}
				}
			}
			g.JSON(400, gin.H{"status": "success", "msg": "ups something went wrong"})

		} else {
			g.JSON(400, gin.H{"status": "success", "msg": "player or queue not found"})
		} 
		
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}