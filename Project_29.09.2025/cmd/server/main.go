// cmd/server/main.go
// Запуск HTTP-сервера и воркеров.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	myhttp "w/internal/http"
	"w/internal/tasks"
)

func main() {
	downloadDir := "downloads"
	os.MkdirAll(downloadDir, 0755)

	manager := tasks.NewTaskManager(100)

	for i := 0; i < 3; i++ {
		tasks.StartWorker(i, manager, downloadDir)
	}

	router := myhttp.NewRouter(manager)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
