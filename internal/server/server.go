package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RalfDalfs/todo-list/internal/api"
	"github.com/RalfDalfs/todo-list/internal/config"
)

func Run(cfg *config.Config) error {
	log.Printf("Servet start port:%s\n", cfg.Port)
	api.Init(cfg)
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
}
