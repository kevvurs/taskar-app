package main

import (
  "bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
			  "REQUEST FORMAT IS NOT VALID")  // 400
			return
		}
    if rs := Event2Record(&de); rs != nil {
      if err := save(rs); err != nil {
      	log.Printf("ERROR: db transaction failed <%v>\n", err)
      	formatter.Text(w, http.StatusInternalServerError,
      	  "SERVER ERROR")  // 500
      	return
      }
    }
		formatter.JSON(w, http.StatusNoContent, nil)  // 204
	}
}

func dataHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		v := vars["filter"]
		t, err := Filter(v)
		if err != nil {
			log.Printf("ERROR: bad query <%v>\n", err)
			formatter.Text(w, http.StatusBadRequest,
			  "QUERY PARAM IS NOT VALID")  // 400
			return
		}
		log.Printf("INFO: searching records since %s", t)
		rs, err := fetch(&t)
		if err != nil {
			log.Printf("ERROR: db search failed <%v>\n", err)
			formatter.Text(w, http.StatusInternalServerError,
		    "SERVER ERROR")  // 500
			return
		}
		if rs == nil {
			log.Println("WARN: no data found")
			formatter.JSON(w, http.StatusOK, []string{})  // 500
			return
		}
		formatter.JSON(w, http.StatusOK, rs)  // 200
	}
}
