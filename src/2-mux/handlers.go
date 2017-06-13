package main

import (
	"data"
	"net/http"
)

func listUsers(w http.ResponseWriter, req *http.Request) {
	// get users
	allUsers := data.GetAll()

	// return users
	sendJSONResponse(allUsers, w)
}
