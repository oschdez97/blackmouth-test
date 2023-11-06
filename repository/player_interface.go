package repository

import "github.com/oschdez97/blackmouth-test/model"

type PlayerRepositoryInterface interface {
	SelectPlayer() []model.Player
	SelectPlayerById(id uint) []model.Player
	InsertPlayer(post model.PostPlayer) []model.Player
	UpdatePlayer(id uint, post model.PostPlayer) []model.Player
	DeletePlayer (id uint) bool
}