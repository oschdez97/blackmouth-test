package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/oschdez97/blackmouth-test/controller"
	"github.com/oschdez97/blackmouth-test/model"
)


func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}


func TestInsertPlayer(t *testing.T) {
	r := SetUpRouter()

	playerController := controller.NewPlayerController(DB)

	r.POST("/api/v1/player", playerController.InsertPlayer)
	player := model.PostPlayer{Name: "John Wick"}

	jsonValue, _ := json.Marshal(player)
	req, _ := http.NewRequest("POST", "/api/v1/player", bytes.NewBuffer(jsonValue))
	
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}



func TestGetPlayer(t *testing.T) {
	r := SetUpRouter()

	playerController := controller.NewPlayerController(DB)

	r.GET("/api/v1/player", playerController.GetPlayer)
	req, _ := http.NewRequest("GET", "/api/v1/player", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var data map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &data)
	
	players := data["data"]	
	
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, players)
}


// Write some more test here...