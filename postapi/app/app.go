package app

import (
	"postapi/app/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/delete", a.DeletePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/update", a.UpdatePutHandler()).Methods("PUT")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("GET")
}
