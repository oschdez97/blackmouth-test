package controller

import (
	"strconv"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/oschdez97/blackmouth-test/model"
	"github.com/oschdez97/blackmouth-test/repository"

	_ "github.com/oschdez97/blackmouth-test/docs"
)

type PlayerController struct {
	DB *sql.DB
}

func NewPlayerController(db *sql.DB) PlayerControllerInterface {
	return &PlayerController{DB: db}
}

// GetPlayers    godoc
// @Summary      Get players array
// @Description  Responds with the list of all players as JSON.
// @Tags         players
// @Produce      json
// @Success      200  {array}  model.Player
// @Router       /player [get]
func (m *PlayerController) GetPlayer(g *gin.Context) {
	db := m.DB
	repo_player := repository.NewPlayerRepository(db)
	get_player := repo_player.SelectPlayer()
	if get_player != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_player, "msg": "get player successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get player successfully"})
	}
}

// GetPlayerById    godoc
// @Summary      	Get player by id
// @Description  	Responds with the player with the given id as JSON.
// @Tags         	players
// @Produce      	json
// @Param        	id  path	string  true	"Player id"
// @Success      	200  {array}  model.Player
// @Router       	/player/{id} [get]
func (m *PlayerController) GetPlayerById(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)

	repo_player := repository.NewPlayerRepository(db)
	get_player := repo_player.SelectPlayerById(uint(id64))
	if get_player != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_player, "msg": "get player successfully"})
	} else {
		g.JSON(404, gin.H{"status": "success", "data": nil, "msg": "player not found"})
	}
}

// PostPlayer    godoc
// @Summary      Store a new player
// @Description  Takes a player JSON and store in DB.
// @Tags         players
// @Produce      json
// @Param        player  body	model.PostPlayer  true  "Player JSON"
// @Success      200  {array}  model.Player
// @Router       /player [post]
func (m *PlayerController) InsertPlayer(g *gin.Context) {
	db := m.DB
	var post model.PostPlayer
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_player := repository.NewPlayerRepository(db)
		insert := repo_player.InsertPlayer(post)
		if insert != nil {
			g.JSON(201, gin.H{"status": "success", "data" : insert, "msg": "insert player successfully"})
		} else {
			g.JSON(500, gin.H{"status": "failed", "msg": "insert player failed"})
		}
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}

// UpdatePlayer    	godoc
// @Summary     	Update a player
// @Description 	Update an existing player in DB.
// @Tags        	players
// @Produce     	json
// @Param        	id  path	string  true	"Player id"
// @Param        	player  body	model.PostPlayer  true  "Player JSON"
// @Success     	200  {array}  model.Player
// @Router      	/player/{id} [put]
func (m *PlayerController) UpdatePlayer(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)
	var post model.PostPlayer
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_player := repository.NewPlayerRepository(db)
		select_player := repo_player.SelectPlayerById(uint(id64))
		
		if select_player != nil {
			updated_player := repo_player.UpdatePlayer(uint(id64), post)
			
			if updated_player != nil {
				g.JSON(200, gin.H{"status": "success", "data": updated_player, "msg": "update player successfully"})
			} else {
				g.JSON(500, gin.H{"status": "success", "msg": "ups something went wrong"})
			} 
		} else {
			g.JSON(404, gin.H{"status": "success", "data": select_player, "msg": "player not found"})	
		}
		
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}

// DeletePlayer    	godoc
// @Summary     	Delete a player
// @Description 	Delete an existing player in DB.
// @Tags        	players
// @Produce     	json
// @Param        	id  path	string  true	"Player id"
// @Success     	200  {object}	bool
// @Router      	/player/{id} [delete] 
func (m *PlayerController) DeletePlayer(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)
	repo_player := repository.NewPlayerRepository(db)
	deleted_player := repo_player.DeletePlayer(uint(id64))
	
	if deleted_player {
		g.JSON(200, gin.H{"status": "success", "data": true, "msg": "delete player successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "ups something went wrong"})
	}
}
