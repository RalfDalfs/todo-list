package main

import (
	"github.com/RalfDalfs/todo-list/internal/db"
	"github.com/RalfDalfs/todo-list/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "./scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatalln(err)
	}
	server.Run()

}
