package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/users"
)

// Login  for
type Login struct {
	Username string
	Password string
}

// Register used for struct Register
type Register struct {
	Username string
	Email    string
	Password string
}

// ErrResponse for
type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	if login["message"] == "ok" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	if register["message"] == "ok" {
		resp := register
		json.NewEncoder(w).Encode(resp)
	} else {
		// resp := ErrResponse{Message: "Invalid Registration Data"}
		resp := register
		json.NewEncoder(w).Encode(resp)
	}
}

// StartAPI used for string API using mux
func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println("Starting API on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
