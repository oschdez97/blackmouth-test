package model

type GameSession struct {
	Id uint `json:"id"`
	Status uint `json:"status"`
	QueueId uint `json:"queue_id"`
	PlayerList []Player `json:"player_list"`
}

type PostGameSession struct {
	QueueId uint `json:"queue_id"`
	PlayerId uint `json:"player_id"`
}


