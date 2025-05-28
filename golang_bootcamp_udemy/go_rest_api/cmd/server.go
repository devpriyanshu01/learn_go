package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "This is root route.")
		w.Write([]byte("This is root route"))
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == http.MethodGet {
			w.Write([]byte("get request"))
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Get method on /students"))
			fmt.Println("Get method on /students")
		case http.MethodPost:
			w.Write([]byte("Post method on /students"))
			fmt.Println("Get method on /students")
		case http.MethodPut:
			w.Write([]byte("Put method on /students"))
			fmt.Println("Put method on /students")
		case http.MethodDelete:
			w.Write([]byte("Delete method on /students"))
			fmt.Println("Delete method on /students")
		}
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("execs route"))
	})

	fmt.Println("server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
