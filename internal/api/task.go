package api

import "net/http"

// taskHandler обрабатывает операции с задачей
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	}
}
