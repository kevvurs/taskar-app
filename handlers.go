package main

import (
  "bytes"
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
    payload = bytes.Trim(payload, "\x00")
		var de DeviceEvent
		if err := json.Unmarshal(payload, &de); err != nil {
		  log.Printf("ERROR: bad request <%v>\n", err)
			formatter.Text(w, http.StatusBadRequest,
			  "REQUEST FORMAT IS NOT VALID")
			return
		}
    if rs := Event2Record(&de); rs != nil {
      if err := upload(rs); err != nil {
      	log.Printf("ERROR: db transaction failed <%v>", err)
      }
    }
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
