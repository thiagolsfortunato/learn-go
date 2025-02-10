package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "login.html", nil)
}
