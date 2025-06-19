package api

import (
	"encoding/json"
	"github.com/RalfDalfs/todo-list/internal/db"
	"net/http"
	"time"
)

// taskHandler обрабатывает операции с задачей
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodPut:
		putTaskHandler(w, r)
	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodDelete:
		deleteHandler(w, r)
	}
}

// deleteHandler обработчик удаления задачи
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		sendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}
	err := db.DeleteTask(id)
	if err != nil {
		sendJSONError(w, "Ошибка удаления задачи"+err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, map[string]string{}, http.StatusOK)
	return
}

// getTaskHandler выдаёт задачу по её айди
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	writeJSON(w, task, http.StatusOK)
}

// putTaskHandler обработчик изменения задачи
func putTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		sendJSONError(w, "Ошибка десериализации JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if task.ID == "" { // Добавить проверку
		sendJSONError(w, "Не указан ID задачи", http.StatusBadRequest)
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

	err = db.UpdateTask(&task)
	if err != nil {
		sendJSONError(w, "Ошибка обновлелния задачи: "+err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}

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

	writeJSON(w, map[string]int64{"id": id}, http.StatusOK)
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusCreated)
	//	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
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
		next, err = nextDate(now, task.Date, task.Repeat)
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
