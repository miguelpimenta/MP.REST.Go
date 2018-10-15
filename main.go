package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var port = 8080
	r := mux.NewRouter()

	println("Running http server @ localhost:" + strconv.Itoa(port))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "REST API in GO!!!")
	})
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", ReadUser).Methods("GET")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", ListAllUsers).Methods("GET")
	if err := http.ListenAndServe(":"+strconv.Itoa(port), r); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet!!!")
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet!!!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet!!!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet!!!")
}

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented yet!!!")
}
