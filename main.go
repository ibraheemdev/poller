package main

import (
	"github.com/ibraheemdev/poller/config"
	"github.com/ibraheemdev/poller/config/db"
	"github.com/ibraheemdev/poller/config/router"
)

func main() {
	config := config.ReadFile()
	client := db.Connect(config.Database)
	defer db.Disconnect(client)
	router.ListenAndServe(config.Server)
}
