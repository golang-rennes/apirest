package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host = "localhost"
	port = 8888
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null
// time repeat.sh 1000 curl -i -X GET localhost:8888/users/1789 >& /dev/null

func main() {
	r := mux.NewRouter()
	// curl -i -X GET localhost:8888/users && echo
	r.HandleFunc("/users", listUsers).Methods(http.MethodGet)

	// curl -i -X GET localhost:8888/users/1789 && echo
	// curl -i -X GET localhost:8888/users/2017 && echo
	// curl -i -X GET localhost:8888/users/2020 && echo

	// curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo

	server := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving API on http://%s…", server)
	log.Fatal(http.ListenAndServe(server, r))
}
