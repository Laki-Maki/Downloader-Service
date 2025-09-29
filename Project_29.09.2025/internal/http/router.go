// internal/http/router.go
package http

import (
	"net/http"
	"w/internal/tasks"
)

func NewRouter(manager *tasks.TaskManager) http.Handler {
	handler := &Handler{Manager: manager}
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", handler.CreateTask) // POST
	mux.HandleFunc("/task", handler.GetTask)     // GET

	return mux
}
