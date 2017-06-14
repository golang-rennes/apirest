package main

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
)

const (
	host       = "localhost"
	port       = 8888
	swaggerDir = "./swagger-ui-2"
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null
// => 5-6ms
// time repeat.sh 1000 curl -i -X GET localhost:8888/users/1789 >& /dev/null
// => 6ms

func main() {
	server := fmt.Sprintf("%s:%d", host, port)

	// initialize go-restful container and filters
	wsContainer := restful.NewContainer()
	wsContainer.Filter(logger)
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// register web services
	userRsc := &userResource{}
	userRsc.register(wsContainer)

	// initialize OpenAPI file and serve Swagger-UI on /docs
	config := openapi.Config{
		WebServices:    wsContainer.RegisteredWebServices(),
		WebServicesURL: "http://" + server,
		APIPath:        "/apidocs.json",
	}
	wsContainer.Add(openapi.NewOpenAPIService(config))
	http.Handle("/docs/", http.StripPrefix("/docs", http.FileServer(http.Dir(swaggerDir+"/dist"))))
	log.Printf("Serving Swagger-UI on http://%s/docs…", server)

	// run
	http.Handle("/", wsContainer)
	log.Printf("Serving API on http://%s and\n                   apidocs.json on http://%s/apidocs.json…", server, server)
	log.Fatal(http.ListenAndServe(server, nil))
}
