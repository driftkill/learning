package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type userData struct {
	UserID   string `json:"id"`
	Password string `json:"pwd"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

type master struct {
	Mpassword string `json:"mpwd"`
}

var userdata []userData
var admin []master

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userd userData
	_ = json.NewDecoder(r.Body).Decode(&userd)
	userdata = append(userdata, userd)
	json.NewEncoder(w).Encode(userd)
	fmt.Fprintln(w, "User registered successfully.")
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range userdata {
		if item.UserID == params["id"] && item.Password == params["pwd"] {
			fmt.Fprintln(w, "Login successfull !!\nWelcome", item.First, item.Last)
			return
		}
		if item.UserID != params["id"] && item.Password == params["pwd"] {
			fmt.Fprintln(w, "Error - user not found")
			return
		}
	}
}

func seeUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range admin {
		if item.Mpassword == params["mpwd"] {
			fmt.Fprintln(w, "List of all users-")
			json.NewEncoder(w).Encode(userdata)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	userdata = append(userdata, userData{UserID: "sam", Password: "samsam", First: "Smarak", Last: "Mishra"})
	admin = append(admin, master{Mpassword: "isvishal"})
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/login/{id},{pwd}", login).Methods("GET")
	r.HandleFunc("/master/{mpwd}", seeUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", r))
}
