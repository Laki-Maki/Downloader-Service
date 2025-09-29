// internal/tasks/worker.go
// Воркеры, которые выполняют загрузку файлов из очереди.

package tasks

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Запуск одного воркера
func StartWorker(id int, m *TaskManager, downloadDir string) {
    go func() {
        for task := range m.queue {
            task.Status = StatusRunning
            task.UpdatedAt = time.Now()

            for _, url := range task.URLs {
                if err := downloadFile(url, downloadDir); err != nil {
                    task.Status = StatusFailed
                    task.Error = err.Error()
                    task.UpdatedAt = time.Now()
                    goto Next // прерываем цикл, задача упала
                }
            }

            // Если все файлы скачались
            task.Status = StatusDone
            task.UpdatedAt = time.Now()

        Next:
        }
    }()
}

// Скачивание одного файла
func downloadFile(url, dir string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to download %s: %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    filename := filepath.Join(dir, filepath.Base(url))
    out, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}
