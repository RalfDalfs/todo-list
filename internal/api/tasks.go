package api

import (
	"github.com/RalfDalfs/todo-list/internal/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// tasksHandler выводит задачи которы есть в БД
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)

}
