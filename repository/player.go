package repository

import (
	"database/sql"
	"log"

	"github.com/oschdez97/blackmouth-test/model"
)

type PlayerRepository struct {
	DB *sql.DB
}

func NewPlayerRepository(db *sql.DB) PlayerRepositoryInterface {
	return &PlayerRepository{DB: db}
}

// InsertPlayer implements PlayerRepositoryInterface
func (m *PlayerRepository) InsertPlayer(post model.PostPlayer) []model.Player {
	lastInsertId := -1
	stmt, err := m.DB.Prepare("INSERT INTO player (name) VALUES ($1) RETURNING id")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	err2 := stmt.QueryRow(post.Name).Scan(&lastInsertId)
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	player := model.Player{Id: uint(lastInsertId), Name: post.Name}	
	return []model.Player{player}
}

// SelectPlayer implements PlayerRepositoryInterface
func (m *PlayerRepository) SelectPlayer() []model.Player {
	var result []model.Player
	rows, err := m.DB.Query("SELECT * FROM player")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id       uint
			name    string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
		} else {
			player := model.Player{Id: id, Name: name}
			result = append(result, player)
		}
	}
	return result
}


// SelectPlayerById implements PlayerRepositoryInterface
func (m *PlayerRepository) SelectPlayerById(id uint) []model.Player {
	stmt, err := m.DB.Prepare("SELECT * FROM player WHERE id=$1 LIMIT 1")
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
	
	for rows.Next() {
		var (
			id       uint
			name    string
		)
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
		} else {
			player := model.Player{Id: id, Name: name}
			return []model.Player{player}
		}
	}
	return nil
}


// UpdatePlayer implements PlayerRepositoryInterface
func (m *PlayerRepository) UpdatePlayer(id uint, post model.PostPlayer) []model.Player {
	stmt, err := m.DB.Prepare("UPDATE player SET name=$1 WHERE id=$2")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(post.Name, id)
	if err2 != nil {
		log.Println(err2)
		return nil
	}
	player := model.Player{Id: uint(id), Name: post.Name}	
	return []model.Player{player}
}


// DeletePlayer implements PlayerRepositoryInterface
func (m *PlayerRepository) DeletePlayer(id uint) bool {
	stmt, err := m.DB.Prepare("DELETE FROM player WHERE id=$1")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}