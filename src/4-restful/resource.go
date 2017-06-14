package main

import restful "github.com/emicklei/go-restful"

type userResource struct{}

func (u *userResource) register(container *restful.Container) {
	ws := new(restful.WebService)
	// global parameters
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// define routes
	// curl -i -X GET localhost:8888/users && echo

	// curl -i -X GET localhost:8888/users/1789 && echo

	// curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo

	container.Add(ws)
}
