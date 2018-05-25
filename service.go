package main

import (
  "net/http"
  "time"
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *http.Server {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return &http.Server{
		Addr:               ":3027",
		Handler:            n,
		ReadTimeout:        10 *time.Second,
		ReadHeaderTimeout:  10 *time.Second,
		WriteTimeout:       10 *time.Second,
		MaxHeaderBytes:     1 << 20,
	}
}

func initRoutes(mx *mux.Router, f *render.Render) {
	mx.HandleFunc("/ping", pingHandler(f)).Methods("GET")
	mx.HandleFunc("/sensor", sensorHandler(f)).Methods("POST")
	mx.HandleFunc("/data", dataHandler(f)).
	  Methods("GET").
	  Queries("filter", "{filter}")
	mx.PathPrefix("/").
	  Handler(http.FileServer(http.Dir("./assets/")))
}
