package main

import (
	"data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type userController struct {
	beego.Controller
}

// ListUsers retrieves all users
func (u *userController) ListUsers() {
	u.Data["json"] = data.GetAll()
	u.ServeJSON()
}

// GetUser retrieves user data
func (u *userController) GetUser() {
	// get ID from path
	id, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	if id < 0 {
		u.Abort(strconv.Itoa(http.StatusNotFound))
	}

	// get user
	user := data.Get(data.Key(id))
	if user == nil {
		u.Abort(strconv.Itoa(http.StatusNotFound))
	}

	// return user
	u.Data["json"] = &user
	u.ServeJSON()
}

// CreateUser creates an user
func (u *userController) CreateUser() {
	// get user data from body
	var user data.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		u.Abort(strconv.Itoa(http.StatusUnprocessableEntity))
	}

	// add user (updates ID)
	data.Add(&user)

	// return user
	u.Data["json"] = &user
	u.ServeJSON()
}
