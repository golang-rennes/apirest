package main

import (
	"data"
	"encoding/json"
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

func createUser(w http.ResponseWriter, req *http.Request) {
	// get user data from body
	var user data.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Request body is not a valid JSON user"))
		return
	}

	// add user (updates ID)
	data.Add(&user)

	// return user
	sendJSONResponse(&user, w)
}
