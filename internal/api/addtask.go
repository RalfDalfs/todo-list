package api

import (
	"encoding/json"
	"github.com/RalfDalfs/todo-list/internal/db"
	"net/http"
	"time"
)

// addTaskHandler обрабатывает добавление новой задачи
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendJSONError(w, "Ошибка десериализации JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		sendJSONError(w, "не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		sendJSONError(w, "Ошибка обработки даты: "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		sendJSONError(w, "Ошибка добавления задачи: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// sendJSONError преобразует сообщение об ошибки в JSON формат и отправляет его
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// checkDate проверяет дату на корректность
func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(TimeFormat)
	}
	t, err := time.Parse(TimeFormat, task.Date)
	if err != nil {
		return err
	}
	// приводим переменную  now и t чтобы их можно было корректно сравнить
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	if t.Before(now) {
		if task.Repeat == "" {
			task.Date = now.Format(TimeFormat)
		} else {
			task.Date = next
		}
	}
	return nil
}
