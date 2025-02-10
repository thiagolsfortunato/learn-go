package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conn()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userEntity, err := repository.SearchByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.CheckPassword(userEntity.Password, user.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userEntity.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
