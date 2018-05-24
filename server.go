package main

import (
  "github.com/gorilla/mux"
  "github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, f *render.Render) {
	mx.HandleFunc("/ping", pingHandler(f)).Methods("GET")
	mx.HandleFunc("/sensor", sensorHandler(f)).Methods("POST")
}
