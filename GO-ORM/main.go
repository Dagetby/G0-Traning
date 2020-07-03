package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func check(err error, text string) {
	if err != nil {
		log.Fatal(text, err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/heroes", AllHeroes).Methods("GET")
	myRouter.HandleFunc("/name{name}/email{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/name{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/name{name}/email{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

	handleRequests()

}
