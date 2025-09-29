// internal/http/handlers.go
// HTTP-эндпоинты: создание задач и получение статуса.

package http

import (
    "encoding/json"
    "net/http"
    "time"

    "w/internal/tasks"
    "w/internal/util"
)

// Handler оборачивает TaskManager для работы через HTTP
type Handler struct {
    Manager *tasks.TaskManager
}

// POST /tasks — создать задачу
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var body struct {
        URLs []string `json:"urls"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    t := &tasks.Task{
        ID:        util.NewID(),          // генерим ID без сторонних либ
        URLs:      body.URLs,
        Status:    tasks.StatusCreated,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    h.Manager.AddTask(t)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(t)
}

// GET /task?id=xxx — получить статус задачи
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id", http.StatusBadRequest)
        return
    }

    task, ok := h.Manager.GetTask(id)
    if !ok {
        http.NotFound(w, r)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}


