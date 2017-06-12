package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = 8888
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null

func main() {
	// curl -i -X GET localhost:8888/users && echo

	server := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving API on http://%s…", server)
	log.Fatal(http.ListenAndServe(server, nil))
}
