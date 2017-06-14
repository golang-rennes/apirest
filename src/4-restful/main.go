package main

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

const (
	host = "localhost"
	port = 8888
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null
// => 5-6ms
// time repeat.sh 1000 curl -i -X GET localhost:8888/users/1789 >& /dev/null
// => 6ms

func main() {
	// initialize go-restful container
	wsContainer := restful.NewContainer()

	// register web services
	userRsc := &userResource{}
	userRsc.register(wsContainer)

	// run
	http.Handle("/", wsContainer)
	server := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving API on http://%s…", server)
	log.Fatal(http.ListenAndServe(server, nil))
}
