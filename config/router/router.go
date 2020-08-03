package router

import (
	"fmt"
	cfg "github.com/ibraheemdev/poller/config"
	"github.com/ibraheemdev/poller/config/router/middleware"
	"github.com/ibraheemdev/poller/polls"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var config cfg.ServerConfig = cfg.Config.Server

type hostSwitch map[string]http.Handler

func (hs hostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Redirect(w, r, "/notfound", 404)
	}
}

// ListenAndServe :
func ListenAndServe() {
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%s", config.Host, config.Port),
			initRoutes()))
}

func initRoutes() hostSwitch {
	apiRouter := httprouter.New()
	apiRouter.POST("/polls", middleware.Cors(polls.Create()))
	apiRouter.GET("/polls/:id", middleware.Cors(polls.Show()))
	apiRouter.PUT("/polls/:id", middleware.Cors(polls.Update()))

	staticRouter := httprouter.New()
	staticRouter.ServeFiles(
		fmt.Sprintf("%s*filepath", config.Static.HomePage),
		http.Dir(config.Static.BuildPath))

	hs := make(hostSwitch)
	hs[fmt.Sprintf("%s%s:%s", config.API.SubDomain, config.Host, config.Port)] = apiRouter
	hs[fmt.Sprintf("%s%s:%s", config.Static.SubDomain, config.Host, config.Port)] = staticRouter
	return hs
}
