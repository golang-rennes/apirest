package main

import (
	"data"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	restful "github.com/emicklei/go-restful"
)

func findUsers(req *restful.Request, resp *restful.Response) {
	// get query parameter
	firstname := strings.TrimSpace(req.QueryParameter("firstname"))

	// get users
	var users []*data.User
	if firstname == "" {
		users = data.GetAll()
	} else {
		users = data.GetByFirstname(firstname)
	}

	// return users
	resp.WriteEntity(users)
}

func getUser(req *restful.Request, resp *restful.Response) {
	// get ID from path
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil || id < 0 {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(fmt.Sprintf("Invalid id \"%s\"", req.PathParameter("id"))))
		return
	}

	// get user
	user := data.Get(data.Key(id))
	if user == nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(fmt.Sprintf("No user found with id %d", id)))
		return
	}

	// return user
	resp.WriteEntity(user)
}

func createUser(req *restful.Request, resp *restful.Response) {
	// get user data from body
	var user data.User
	err := req.ReadEntity(&user)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
		resp.Write([]byte("Request body is not a valid JSON user"))
		return
	}

	// add user (updates ID)
	data.Add(&user)

	// return user
	resp.WriteEntity(user)
}
