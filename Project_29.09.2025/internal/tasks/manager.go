// internal/tasks/manager.go
// Хранит все задачи и очередь для воркеров.

package tasks

import "sync"

type TaskManager struct {
	mu    sync.Mutex
	tasks map[string]*Task // все задачи по ID
	queue chan *Task       // очередь задач для воркеров
}

// Конструктор
func NewTaskManager(queueSize int) *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*Task),
		queue: make(chan *Task, queueSize),
	}
}

// Добавление задачи в память и очередь
func (m *TaskManager) AddTask(t *Task) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks[t.ID] = t
	m.queue <- t
}

// Получение задачи по ID
func (m *TaskManager) GetTask(id string) (*Task, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	t, ok := m.tasks[id]
	return t, ok
}

// Список всех задач
func (m *TaskManager) AllTasks() []*Task {
	m.mu.Lock()
	defer m.mu.Unlock()
	list := make([]*Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		list = append(list, t)
	}
	return list
}
