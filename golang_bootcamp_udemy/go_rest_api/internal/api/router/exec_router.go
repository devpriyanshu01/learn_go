package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func execRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/execs/", handlers.ExecHandler)
	return mux
}
