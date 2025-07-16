package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func studentRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// mux.HandleFunc("/students/", handlers.StudentHandler)
	mux.HandleFunc("POST /student/", handlers.AddOneStudent)
	mux.HandleFunc("POST /students/", handlers.AddStudents)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudent)

	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudent)
	mux.HandleFunc("GET /students/", handlers.GetStudents)
	mux.HandleFunc("PATCH /students/", handlers.UpdateOneStudent)

	return mux
}
