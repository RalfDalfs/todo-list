package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Run() error {
	godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("TODO_PORT"))
	if err != nil {
		log.Printf("Произошла ошибка преобразования: %v, сервер запущен на стандартном порту 7540", err)
	}
	if port == 0 {
		port = 7540
	}
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
