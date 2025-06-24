package main

import (
	"log"

	"github.com/RalfDalfs/todo-list/internal/config"
	"github.com/RalfDalfs/todo-list/internal/db"
	"github.com/RalfDalfs/todo-list/internal/server"
)

func main() {
	cfg := config.New()
	// Инициализация базы данных
	err := db.Init(cfg.DbFile)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()
	//запуск сервера
	server.Run(cfg)
}
