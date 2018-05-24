package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/unrolled/render"
)

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNoContent, nil)  // 204
	}
}

func sensorHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var sd []SensorPacket
		if err := json.Unmarshal(payload, &sd); err != nil {
		  log.Printf("ERROR: bad request <%v>\n", err)
			formatter.Text(w, http.StatusBadRequest,
			  "REQUEST FORMAT IS NOT VALID")
			return
		}
		log.Printf("INFO: event- %v\n", sd)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
