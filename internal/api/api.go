package api

import (
	"net/http"

	"github.com/RalfDalfs/todo-list/internal/config"
)

// Init инициализирует апи хендлеры
func Init(cfg *config.Config) {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", auth(taskHandler, cfg))
	http.HandleFunc("/api/tasks", auth(tasksHandler, cfg))
	http.HandleFunc("/api/task/done", auth(doneHandler, cfg))
	http.HandleFunc("/api/signin", signinHandler(cfg))
}
