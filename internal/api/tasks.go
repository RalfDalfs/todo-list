package api

import (
	"github.com/RalfDalfs/todo-list/internal/db"
	"net/http"
	"time"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

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
			tasks, err := db.TasksSearch(search, 50)
			if err != nil {
				sendJSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
			return
		}
		tasks, err := db.TasksDate(date, 50)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
	}

}
