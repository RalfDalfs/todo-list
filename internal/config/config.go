package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbFile string //TODO_DBFILE
	Port   string
	Passw  string
}

func New() *Config {
	godotenv.Load()

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "./scheduler.db"
	}

	port, ok := os.LookupEnv("TODO_PORT")
	if !ok || port == "" {
		port = "7540"
	}

	pass := os.Getenv("TODO_PASSWORD")

	return &Config{
		DbFile: dbFile,
		Port:   port,
		Passw:  pass,
	}
}
