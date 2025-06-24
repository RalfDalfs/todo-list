package api

import (
	"net/http"
	"time"

	"github.com/RalfDalfs/todo-list/internal/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

const limit = 50

// tasksHandler выводит задачи которы есть в БД
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	if search == "" {
		tasks, err := db.Tasks(50) // в параметре максимальное количество записей
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
		return
	} else {
		date, err := time.Parse("02.01.2006", search)
		if err != nil {
			tasks, err := db.TasksSearch(search, limit)
			if err != nil {
				sendJSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
			return
		}
		tasks, err := db.TasksDate(date, limit)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
	}

}
