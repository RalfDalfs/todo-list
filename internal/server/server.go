package server

import (
	"fmt"
	"github.com/RalfDalfs/todo-list/internal/api"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func Run() error {
	godotenv.Load()
	port, ok := os.LookupEnv("TODO_PORT")
	if !ok || port == "" {
		port = "7540"
	}

	api.Init()
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
