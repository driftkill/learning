package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Detail comment
type Detail struct {
	ID       string `json:"id"`
	Position string `json:"position"`
	Dept     string `json:"dept"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

var details []Detail

func getDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range details {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Detail{})
}

func createDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var detail Detail
	_ = json.NewDecoder(r.Body).Decode(&detail)
	detail.ID = strconv.Itoa(rand.Intn(100))
	details = append(details, detail)
	json.NewEncoder(w).Encode(detail)
}

func updateDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range details {
		if item.ID == params["id"] {
			details = append(details[:index], details[index+1:]...)
			var detail Detail
			_ = json.NewDecoder(r.Body).Decode(&detail)
			detail.ID = strconv.Itoa(rand.Intn(100))
			details = append(details, detail)
			json.NewEncoder(w).Encode(detail)
		}
	}
	json.NewEncoder(w).Encode(details)
}

func deleteDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range details {
		if item.ID == params["id"] {
			details = append(details[:index], details[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(details)
}

func main() {
	r := mux.NewRouter()
	details = append(details, Detail{ID: "10", Position: "Engg.", Dept: "IT", First: "Vishal", Last: "Kumar"})
	r.HandleFunc("/empdetails", getDetails).Methods("GET")
	r.HandleFunc("/empdetails/{id}", getDetail).Methods("GET")
	r.HandleFunc("/empdetails", createDetail).Methods("POST")
	r.HandleFunc("/empdetails/{id}", updateDetail).Methods("PUT")
	r.HandleFunc("/empdetails/{id}", deleteDetail).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
