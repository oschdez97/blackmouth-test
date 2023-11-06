package repository

import "github.com/oschdez97/blackmouth-test/model"

type GameSessionRepositoryInterface interface {
	SelectGameSession() []model.GameSession
	SelectGameSessionById(id uint) []model.GameSession
	SelectGameSessionByStatus(status uint) []model.GameSession
	InsertGameSession(post model.PostGameSession) int
	UpdateGameSession(gamesession_id int, queue_capacity int, post model.PostGameSession) bool
	FindExistingGameSession(post model.PostGameSession) []model.GameSession
}