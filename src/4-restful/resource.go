package main

import (
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
		Metadata(openapi.KeyOpenAPITags, tags))

	// curl -i -X GET localhost:8888/users/1789 && echo
	ws.Route(ws.GET("/{id}").To(getUser).
		Metadata(openapi.KeyOpenAPITags, tags))

	// curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo
	ws.Route(ws.POST("/").To(createUser).
		Metadata(openapi.KeyOpenAPITags, tags))

	container.Add(ws)
}
