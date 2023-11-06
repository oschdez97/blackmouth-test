package repository

import (
	"database/sql"
	"log"

	"github.com/oschdez97/blackmouth-test/model"
)

type GameSessionRepository struct {
	DB *sql.DB
}

func NewGameSessionRepository(db *sql.DB) GameSessionRepositoryInterface {
	return &GameSessionRepository{DB: db}
}

func getPlayerList(DB *sql.DB, id uint) []model.Player{
	// Get player list
	var players []model.Player

	stmt, err2 := DB.Prepare("SELECT player_id, (SELECT name FROM player WHERE id=gamesession_players.player_id) AS player_name FROM gamesession_players WHERE gamesession_id=$1")
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	defer stmt.Close()
	rows, err3 := stmt.Query(id)
	if err3 != nil {
		log.Println(err3)
		return nil
	}

	for rows.Next() {
		var (
			id   uint
			name string
		)
		err4 := rows.Scan(&id, &name)
		if err4 != nil {
			log.Println(err4)
		} else {
			player := model.Player{Id: id, Name: name}
			players = append(players, player)
		}
	}

	return players
}

func (m *GameSessionRepository) SelectGameSession() []model.GameSession {
	var result []model.GameSession

	rows, err := m.DB.Query("SELECT * FROM gamesession")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id       uint
			status   uint
			queue_id uint
		)
		err := rows.Scan(&id, &status, &queue_id)
		if err != nil {
			log.Println(err)
		} else {
			players := getPlayerList(m.DB, id)
			gamesession := model.GameSession{Id: id, Status: status, QueueId: queue_id, PlayerList: players}
			result = append(result, gamesession)
		}
	}
	return result
}

func (m *GameSessionRepository) SelectGameSessionById(id uint) []model.GameSession {
	stmt, err := m.DB.Prepare("SELECT * FROM gamesession WHERE id=$1 LIMIT 1")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(id)
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	var result []model.GameSession
	
	for rows.Next() {
		var (
			id       uint
			status   uint
			queue_id uint
		)
		err := rows.Scan(&id, &status, &queue_id)
		if err != nil {
			log.Println(err)
		} else {
			players := getPlayerList(m.DB, id)
			gamesession := model.GameSession{Id: id, Status: status, QueueId: queue_id, PlayerList: players}
			result = append(result, gamesession)
		}
	}
	return result
	
}

func (m *GameSessionRepository) SelectGameSessionByStatus(status uint) []model.GameSession {
	stmt, err := m.DB.Prepare("SELECT * FROM gamesession WHERE status=$1")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(status)
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	var result []model.GameSession
	
	for rows.Next() {
		var (
			id       uint
			status   uint
			queue_id uint
		)
		err := rows.Scan(&id, &status, &queue_id)
		if err != nil {
			log.Println(err)
		} else {
			players := getPlayerList(m.DB, id)
			gamesession := model.GameSession{Id: id, Status: status, QueueId: queue_id, PlayerList: players}
			result = append(result, gamesession)
		}
	}
	return result
	
}

func (m *GameSessionRepository) FindExistingGameSession(post model.PostGameSession) []model.GameSession {	
	var result []model.GameSession
	
	stmt, err := m.DB.Prepare("SELECT * FROM gamesession WHERE status=1 AND queue_id=$1")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(post.QueueId)
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	
	for rows.Next() {
		var (
			id        uint
			status 	  uint
			queue_id  uint
		)
		err := rows.Scan(&id, &queue_id, &status)
		if err != nil {
			log.Println(err)
		} else {
			gamesession := model.GameSession{Id: id, Status: status, QueueId: queue_id}			
			result = append(result, gamesession)
			return result
		}
	}
	return nil
}

func (m *GameSessionRepository) InsertGameSession(post model.PostGameSession) int {
	lastInsertId := -1
	
	stmt, err := m.DB.Prepare("INSERT INTO gamesession (status, queue_id) VALUES ($1, $2) RETURNING id")
	if err != nil {
		log.Println(err)
		return -1
	}
	defer stmt.Close()
	err2 := stmt.QueryRow(1, post.QueueId).Scan(&lastInsertId)
	if err2 != nil {
		log.Println(err2)
		return -1
	}
	return lastInsertId
}

func (m *GameSessionRepository) UpdateGameSession(gamesession_id int, queue_capacity int, post model.PostGameSession) bool {
	stmt, err := m.DB.Prepare("INSERT INTO gamesession_players (gamesession_id, player_id) VALUES ($1, $2)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(gamesession_id, post.PlayerId)
	if err2 != nil {
		log.Println(err2)
		return false
	}

	// check if the game session's queue is full
	stmt, err3 := m.DB.Prepare("SELECT COUNT(id) AS total FROM gamesession_players WHERE gamesession_id=$1")
	if err3 != nil {
		log.Println(err3)
		return false
	}
	defer stmt.Close()
	rows, err4 := stmt.Query(gamesession_id)
	if err4 != nil {
		log.Println(err4)
		return false
	}

	for rows.Next() {
		var (
			total	int
		)
		err5 := rows.Scan(&total)
		if err5 != nil {
			log.Println(err5)
		} else {
			
			if total >= queue_capacity {
				// close the game session

				stmt, err6 := m.DB.Prepare("UPDATE gamesession SET status=0 WHERE id=$1")
				if err6 != nil {
					log.Println(err6)
					return false
				}
				defer stmt.Close()
				_, err7 := stmt.Exec(gamesession_id)
				if err7 != nil {
					log.Println(err7)
					return false
				}

			}
		}
	}
	return true
}