package repository

import "github.com/oschdez97/blackmouth-test/model"

type QueueRepositoryInterface interface {
	SelectQueue() []model.Queue
	SelectQueueById(id uint) []model.Queue
	InsertQueue(post model.PostQueue, queue_count int) []model.Queue
	UpdateQueue(id uint, post model.PutQueue) []model.Queue
	DeleteQueue (id uint) bool
}