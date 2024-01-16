package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"crud/database"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Error when read request body"))
		return
	}

	var user user

	if error := json.Unmarshal(requestBody, &user); error != nil {
		w.Write([]byte("Error when serialize user"))
		return
	}

	fmt.Println(user)

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to Database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("insert into users (name, email) values (? , ?)")
	if error != nil {
		w.Write([]byte("Error creating the db statement"))
		return
	}
	defer statement.Close()

	insert, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Error when execute statement"))
		return
	}

	lastId, error := insert.LastInsertId()
	if error != nil {
		w.Write([]byte("Error getting the entered ID"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted with sucess! ID: %d", lastId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to Database"))
		return
	}
	defer db.Close()

	lines, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Error getting users"))
		return
	}
	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if error := lines.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error getting the user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if error := json.NewEncoder(w).Encode(users); error != nil {
		w.Write([]byte("Error converting user to json"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to Int"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to Database"))
		return
	}
	defer db.Close()

	line, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		w.Write([]byte("Error getting user"))
		return
	}
	defer line.Close()

	var user user
	if line.Next() {
		if error := line.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error getting the user"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		if error := json.NewEncoder(w).Encode(user); error != nil {
			w.Write([]byte("Error converting user to json"))
			return
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to Int"))
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Error when read request body"))
		return
	}

	var user user

	if error := json.Unmarshal(requestBody, &user); error != nil {
		w.Write([]byte("Error serializing user"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to Database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update users set name = ?, email = ? where id = ?")
	if error != nil {
		w.Write([]byte("Error creating the db statement"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Email, ID); error != nil {
		w.Write([]byte("Error updating the user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		w.Write([]byte("Error converting parameter to Int"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to Database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("delete from users where id = ?")
	if error != nil {
		w.Write([]byte("Error creating the db statement"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("Error deleting the user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
