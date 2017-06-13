package main

import (
	"data"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func listUsers(w http.ResponseWriter, req *http.Request) {
	// get users
	allUsers := data.GetAll()

	// return users
	sendJSONResponse(allUsers, w)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	// get ID from path
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if id < 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Invalid id \"%s\"", vars["id"])))
		return
	}

	// get user
	user := data.Get(data.Key(id))
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No user found with id %d", id)))
		return
	}

	// return user
	sendJSONResponse(user, w)
}
