package main

import (
	"data"
	"net/http"

	restful "github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
)

type userResource struct{}

func (u *userResource) register(container *restful.Container) {
	ws := new(restful.WebService)
	// global parameters
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"users"}

	// define routes
	// curl -i -X GET localhost:8888/users && echo
	ws.Route(ws.GET("/").To(findUsers).
		Metadata(openapi.KeyOpenAPITags, tags).
		Doc("gets all users list").
		Param(ws.QueryParameter("firstname", "filter users with this firstname").DataType("string").Required(false)).
		Writes([]data.User{}).
		Returns(http.StatusOK, "success, users list is returned", []data.User{}))

	// curl -i -X GET localhost:8888/users/1789 && echo
	ws.Route(ws.GET("/{id}").To(getUser).
		Metadata(openapi.KeyOpenAPITags, tags).
		Doc("gets a specific user").
		Param(ws.PathParameter("id", "identifier of the user").DataType("integer").Required(true)).
		Writes(data.User{}).
		Returns(http.StatusOK, "success, user is returned", data.User{}).
		Returns(http.StatusNotFound, "user does not exist", nil))

	// curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo
	ws.Route(ws.POST("/").To(createUser).
		Metadata(openapi.KeyOpenAPITags, tags).
		Doc("creates a user").
		Reads(data.User{}).
		Writes(data.User{}).
		Returns(http.StatusOK, "success, created user is returned", data.User{}))

	container.Add(ws)
}
