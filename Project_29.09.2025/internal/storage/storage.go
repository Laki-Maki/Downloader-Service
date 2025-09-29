package storage

import "w/internal/tasks"

type Storage struct {
	tasks map[string]tasks.Task
}

func NewStorage() *Storage {
	return &Storage{
		tasks: make(map[string]tasks.Task),
	}
}
