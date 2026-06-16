package main

import (
	"log"

	"github.com/priyanshu334/gopr1/internal/config"
	"github.com/priyanshu334/gopr1/internal/database"
	"github.com/priyanshu334/gopr1/internal/server"
	"github.com/priyanshu334/gopr1/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	logg, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	srv := server.New(cfg, logg, db)
	logg.Info("starting server on port " + cfg.AppPort)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
