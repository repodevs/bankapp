package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/repodevs/bankapp/helpers"
	"github.com/repodevs/bankapp/interfaces"
	"github.com/repodevs/bankapp/users"
)

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody interfaces.Login
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var formattedBody interfaces.Register
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	apiResponse(register, w)
}

// readBody from http.Request
func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)
	return body
}

// apiResponse used for response
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "ok" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := interfaces.ErrResponse{Message: "invalid data"}
		// resp := call
		json.NewEncoder(w).Encode(resp)
	}
}

// StartAPI used for starting API using mux
func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println("Starting API on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
