package handlers

import (
	"fmt"
	"net/http"
)

func ExecHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Write([]byte("inside /exec - post method"))
		query := r.URL.Query()
		fmt.Println("Query Received:", query)
		fmt.Println("name", query.Get("name"))
		fmt.Println("name", query.Get("age"))
	}
}
