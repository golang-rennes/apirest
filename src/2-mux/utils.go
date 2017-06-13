package main

import (
	"encoding/json"
	"net/http"
)

func sendJSONResponse(v interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
