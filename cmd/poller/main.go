package main

import (
	_ "github.com/ibraheemdev/poller/config"
	"github.com/ibraheemdev/poller/internal/router"
	"github.com/ibraheemdev/poller/pkg/database"
)

func main() {
	client := database.Connect()
	defer database.Disconnect(client)
	router.ListenAndServe()
}
