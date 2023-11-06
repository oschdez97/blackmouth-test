package model

type GameSessionPlayers struct {
	Id uint `json:"id"`
	GameSessionId uint `json:"gamesession_id"`
	PlayerId uint `json:"player_id"`
}

