// internal/tasks/task.go
// Определение структуры Task и статусов задачи.

package tasks

import "time"

// Возможные статусы задачи
type Status string

const (
    StatusCreated Status = "created"
    StatusRunning Status = "in_progress"
    StatusDone    Status = "done"
    StatusFailed  Status = "failed"
)

// Task — одна задача скачивания
type Task struct {
    ID        string    `json:"id"`          // уникальный идентификатор
    URLs      []string  `json:"urls"`        // список ссылок для скачивания
    Status    Status    `json:"status"`      // текущий статус
    CreatedAt time.Time `json:"created_at"`  // время создания
    UpdatedAt time.Time `json:"updated_at"`  // время последнего обновления
    Error     string    `json:"error,omitempty"` // ошибка, если есть
}
