package controller

import (
	"strconv"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/oschdez97/blackmouth-test/model"
	"github.com/oschdez97/blackmouth-test/repository"
)

type QueueController struct {
	DB *sql.DB
	QueueCount int
}

func NewQueueController(db *sql.DB, queue_count int) QueueControllerInterface {
	return &QueueController{DB: db, QueueCount: queue_count}
}

// GetQueues	 godoc
// @Summary      Get queues array
// @Description  Responds with the list of all queues as JSON.
// @Tags         queues
// @Produce      json
// @Success      200  {array}  model.Queue
// @Router       /queue [get]
func (m *QueueController) GetQueue(g *gin.Context) {
	db := m.DB

	repo_queue := repository.NewQueueRepository(db)
	get_queue := repo_queue.SelectQueue()
	if get_queue != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_queue, "msg": "get queue successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get queue successfully"})
	}
}

// PostQueue    godoc
// @Summary      Store a new queue
// @Description  Takes a queue JSON and store in DB.
// @Tags         queues
// @Produce      json
// @Param        queue  body	model.PostQueue  true  "Queue JSON"
// @Success      200  {array}  model.Queue
// @Router       /queue [post]
func (m *QueueController) InsertQueue(g *gin.Context) {
	db := m.DB
	queue_count := m.QueueCount
	var post model.PostQueue
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_queue := repository.NewQueueRepository(db)
		insert := repo_queue.InsertQueue(post, queue_count)
		if insert != nil {
			g.JSON(200, gin.H{"status": "success", "data" : insert, "msg": "insert queue successfully"})
		} else {
			g.JSON(500, gin.H{"status": "failed", "msg": "insert queue failed"})
		}
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}

// GetQueueById    godoc
// @Summary      	Get queue by id
// @Description  	Responds with the queue with the given id as JSON.
// @Tags         	queues
// @Produce      	json
// @Param        	id  path	string  true	"Queue id"
// @Success      	200  {array}  model.Queue
// @Router       	/queue/{id} [get]
func (m *QueueController) GetQueueById(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)

	repo_queue := repository.NewQueueRepository(db)
	get_queue := repo_queue.SelectQueueById(uint(id64))
	if get_queue != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_queue, "msg": "get queue successfully"})
	} else {
		g.JSON(404, gin.H{"status": "success", "data": nil, "msg": "queue not found"})
	}
}

// UpdateQueue    	godoc
// @Summary     	Update a queue
// @Description 	Update an existing queue in DB.
// @Tags        	queues
// @Produce     	json
// @Param        	id  path	string  true	"Queue id"
// @Param        	queue  body	model.PutQueue  true  "Queue JSON"
// @Success     	200  {array}  model.Queue
// @Router      	/queue/{id} [put]
func (m *QueueController) UpdateQueue(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)
	var post model.PutQueue
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_queue := repository.NewQueueRepository(db)
		select_queue := repo_queue.SelectQueueById(uint(id64))
		
		if select_queue != nil {
			updated_queue := repo_queue.UpdateQueue(uint(id64), post)
			
			if updated_queue != nil {
				g.JSON(200, gin.H{"status": "success", "data": updated_queue, "msg": "update queue successfully"})
			} else {
				g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "ups something went wrong"})
			}
		} else {
			g.JSON(404, gin.H{"status": "success", "msg": "queue not found"})
		}
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}

// DeleteQueue    	godoc
// @Summary     	Delete a queue
// @Description 	Delete an existing queue in DB.
// @Tags        	queues
// @Produce     	json
// @Param        	id  path	string  true	"Queue id"
// @Success     	200  {object}	bool
// @Router      	/queue/{id} [delete] 
func (m *QueueController) DeleteQueue(g *gin.Context) {
	db := m.DB
	id64, _ := strconv.ParseUint(g.Param("id"), 10, 32)
	repo_queue := repository.NewQueueRepository(db)
	deleted_queue := repo_queue.DeleteQueue(uint(id64))
	
	if deleted_queue {
		g.JSON(200, gin.H{"status": "success", "data": true, "msg": "delete queue successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "ups something went wrong"})
	}
}
