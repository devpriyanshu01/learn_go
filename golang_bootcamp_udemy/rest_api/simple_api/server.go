package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Handling incoming orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Handling users...")
	})

	port := 3000
	fmt.Printf("Server is running on port: %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}