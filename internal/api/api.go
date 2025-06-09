package api

import (
	"net/http"
	"time"
)

// Init инициализирует апи хендлеры
func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
	http.HandleFunc("/api/task", taskHandler)
}

// nextDayHandler возвращает следующую дату для задачи
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	if now == "" {
		now = time.Now().Format(TimeFormat)
	}
	nowForm, err := time.Parse(TimeFormat, now)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	nextDate, err := NextDate(nowForm, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}

// taskHandler обрабатывает операции с задачами
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}
