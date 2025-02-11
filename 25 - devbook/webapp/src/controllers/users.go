package controllers

import (
	"webapp/src/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
)

// createUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf("%s%s", config.ApiUrl, "/users"), "application/json", bytes.NewBuffer(user))
 
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	
	responses.JSON(w, res.StatusCode, nil)
}
