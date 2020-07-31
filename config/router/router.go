package router

import (
	"github.com/ibraheemdev/poller/config/router/middleware"
	"github.com/ibraheemdev/poller/polls"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// ListenAndServe :
func ListenAndServe(config struct {
	Port string "yaml:\"port\""
	Host string "yaml:\"host\""
}) {
	log.Fatal(http.ListenAndServe(config.Host+config.Port, initRoutes()))
}

func initRoutes() *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/polls", middleware.Cors(polls.Create()))
	router.GET("/api/polls/:id", polls.Show())
	router.PUT("/api/polls/:id", middleware.Cors(polls.Update()))
	router.ServeFiles("/static/*filepath", http.Dir("client/build"))
	return router
}
