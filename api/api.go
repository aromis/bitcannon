package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
)

type API struct {
	M *martini.ClassicMartini
}

func NewAPI() *API {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	return &API{m}
}

func (api *API) AddRoutes() {
	api.M.Get("/stats", torrentDB.Count)
	api.M.Get("/browse", torrentDB.Categories)
	api.M.Get("/browse/:category", torrentDB.Browse)
	api.M.Get("/torrent/:btih", torrentDB.Get)
	api.M.Get("/search/:query", torrentDB.Search)
	api.M.Get("/search/:query/:page", torrentDB.Search)
}

func (api *API) Run(port string) {
	api.M.RunOnAddr(port)
}
