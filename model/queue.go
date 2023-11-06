package model

type Queue struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Capacity uint `json:"capacity"`
}

type PostQueue struct {
	Name string `json:"name" binding:"required"`
	Capacity uint `json:"capacity" binding:"required"`
}

type PutQueue struct {
	Name string `json:"name" binding:"required"`
}