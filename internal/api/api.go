package api

import (
	"github.com/RalfDalfs/todo-list/internal/nextdate"
	"net/http"
	"time"
)

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	if now == "" {
		now = time.Now().Format(nextdate.TimeFormat)
	}
	nowForm, err := time.Parse(nextdate.TimeFormat, now)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	nextDate, err := nextdate.NextDate(nowForm, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
