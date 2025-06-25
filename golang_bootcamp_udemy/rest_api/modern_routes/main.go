package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//new way of routing
	mux.HandleFunc("POST /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item created")
	})

	mux.HandleFunc("DELETE /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item deleted")
	})

	//new way to receive path parameter
	mux.HandleFunc("GET /product/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Item of the product is: %s", id)
	})

	//below handler will throw error.
	// mux.HandleFunc("GET /{param}/buy", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Param is: %s", r.PathValue("param"))
	// })

	//for below handler, if you send /product/id explicitely then below handler will called else
	//the handler before the commented part will be called. This is the relation b/n /product/{id} &
	//                                                                               /prodct/id.
	mux.HandleFunc("GET /product/id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "nothnig to display here...")
	})

	fmt.Println("SERVER IS RUNNING ON PORT - 8080")
	http.ListenAndServe(":8080", mux)
}
