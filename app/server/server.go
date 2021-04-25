package server

import (
	"log"
	"ml-challenge/config"
)

func Init() {
	cfg := config.GetConfig()
	r := NewRouter()
	log.Fatal(r.Run(cfg.GetString("server.port")))
}
