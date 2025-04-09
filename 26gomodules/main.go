package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learn Modules in Golang")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") 	

	log.Fatal(http.ListenAndServe(":3001", r))
}

func Greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang Series on YT</h1>"))
}
