package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "login.html", nil)
}

// LoadCreateUserPage loads the Create User page
func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "create-user.html", nil)
}
