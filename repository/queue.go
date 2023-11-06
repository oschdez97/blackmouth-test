package repository

import (
	"database/sql"
	"log"

	"github.com/oschdez97/blackmouth-test/model"
)

type QueueRepository struct {
	DB *sql.DB
}

func NewQueueRepository(db *sql.DB) QueueRepositoryInterface {
	return &QueueRepository{DB: db}
}

// InsertQueue implements QueueRepositoryInterface
func (m *QueueRepository) InsertQueue(post model.PostQueue, queue_count int) []model.Queue {
	rows, err := m.DB.Query("SELECT COUNT(id) AS total FROM queue")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			total	int
		)
		err := rows.Scan(&total)
		if err != nil {
			log.Println(err)
		} else {
			if total >= queue_count {
				return nil
			} else {
				lastInsertId := -1
				stmt, err := m.DB.Prepare("INSERT INTO queue (name, capacity) VALUES ($1, $2) RETURNING id")
				if err != nil {
					log.Println(err)
					return nil
				}
				defer stmt.Close()
				err2 := stmt.QueryRow(post.Name, post.Capacity).Scan(&lastInsertId)
				if err2 != nil {
					log.Println(err2)
					return nil
				}
				queue := model.Queue{Id: uint(lastInsertId), Name: post.Name, Capacity: post.Capacity }	
				return []model.Queue{queue}
			}
		}
	}
	return nil
}

// SelectQueue implements QueueRepositoryInterface
func (m *QueueRepository) SelectQueue() []model.Queue {
	var result []model.Queue
	rows, err := m.DB.Query("SELECT * FROM queue")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id       uint
			name	 string
			capacity uint
		)
		err := rows.Scan(&id, &name, &capacity)
		if err != nil {
			log.Println(err)
		} else {
			queue := model.Queue{Id: id, Name: name, Capacity: capacity}
			result = append(result, queue)
		}
	}
	return result
}

// SelectQueueById implements QueueRepositoryInterface
func (m *QueueRepository) SelectQueueById(id uint) []model.Queue {
	stmt, err := m.DB.Prepare("SELECT * FROM queue WHERE id=$1 LIMIT 1")
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
			name	 string
			capacity uint
		)
		err := rows.Scan(&id, &name, &capacity)
		if err != nil {
			log.Println(err)
		} else {
			queue := model.Queue{Id: id, Name: name, Capacity: capacity}
			return []model.Queue{queue}
		}
	}
	return nil
}


// UpdateQueue implements QueueRepositoryInterface
func (m *QueueRepository) UpdateQueue(id uint, post model.PutQueue) []model.Queue {
	stmt, err := m.DB.Prepare("UPDATE queue SET name=$1 WHERE id=$2")
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
	queue := model.Queue{Id: id, Name: post.Name }	
	return []model.Queue{queue}
}


// DeleteQueue implements QueueRepositoryInterface
func (m *QueueRepository) DeleteQueue(id uint) bool {
	stmt, err := m.DB.Prepare("DELETE FROM queue WHERE id=$1")
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