package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	host = "localhost"
	port = 8888
)

// time repeat.sh 1000 curl -i -X GET localhost:8888/users >& /dev/null
// => 5-6ms
// time repeat.sh 1000 curl -i -X GET localhost:8888/users/1789 >& /dev/null
// => < 5ms

func main() {
	// config
	beego.ErrorHandler(strconv.Itoa(http.StatusUnprocessableEntity), unprocessableEntity)

	// curl -i -X GET localhost:8888/users && echo
	// curl -i -X POST localhost:8888/users -d '{"firstname":"toto", "lastname":"titi"}' && echo
	beego.Router("/users", &userController{}, "get:ListUsers;post:CreateUser")

	// curl -i -X GET localhost:8888/users/1789 && echo
	// curl -i -X GET localhost:8888/users/2017 && echo
	// curl -i -X GET localhost:8888/users/2020 && echo
	beego.Router("/users/:id:int", &userController{}, "get:GetUser")

	server := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Serving API on http://%s…", server)
	beego.Run(server)
}

func unprocessableEntity(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte("Request body is not valid JSON or contains unexpected fields"))
}
