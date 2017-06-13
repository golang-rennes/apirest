package main

import (
	"data"
	"net/http"
)

func listUsers(w http.ResponseWriter, req *http.Request) {
	// request filtering
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get users
	allUsers := data.GetAll()

	// return users
	sendJSONResponse(allUsers, w)
}
