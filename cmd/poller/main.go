package main

import (
	_ "github.com/ibraheemdev/poller/config"
	"github.com/ibraheemdev/poller/config/db"
	"github.com/ibraheemdev/poller/config/router"
)

func main() {
	client := db.Connect()
	defer db.Disconnect(client)
	router.ListenAndServe()
}
