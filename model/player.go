package model

type Player struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

type PostPlayer struct {
	Name string `json:"name" binding:"required"`
}