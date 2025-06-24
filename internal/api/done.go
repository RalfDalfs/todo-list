package api

import (
	"net/http"
	"time"

	"github.com/RalfDalfs/todo-list/internal/db"
)

// doneHandler делает задачу выполненной если нет правила рипит и удаляет её из бд если рипит есть ставит след. дату
func doneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendJSONError(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		sendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		sendJSONError(w, "Ошибка получения задачи: "+err.Error(), http.StatusBadRequest)
		return
	}
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			sendJSONError(w, "Ошибка удаления задачи"+err.Error(), http.StatusBadRequest)
			return
		}
		writeJSON(w, map[string]string{}, http.StatusOK)
		return
	} else {
		now := time.Now().Format(TimeFormat)
		nowForm, err := time.Parse(TimeFormat, now)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		next, err := nextDate(nowForm, task.Date, task.Repeat)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = db.UpdateDate(next, id)
		if err != nil {
			sendJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeJSON(w, map[string]string{}, http.StatusOK)
		return
	}

}
