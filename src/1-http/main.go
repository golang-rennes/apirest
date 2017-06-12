package main

import (
	"data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = 8888
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null
// => < 5ms

func main() {
	// curl -i -X GET localhost:8888/users && echo
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		// request filtering
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// get users
		allUsers := data.GetAll()

		// return users
		js, err := json.Marshal(allUsers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	server := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving API on http://%s…", server)
	log.Fatal(http.ListenAndServe(server, nil))
}
