package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("GET /teachers/", handlers.GetTeacherHandler)
	mux.HandleFunc("POST /teachers/", handlers.AddTeacherHandler)
	mux.HandleFunc("PUT /teachers/", handlers.UpdateTeacher)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.UpdateTeacherFieldsPatch)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacher)
	mux.HandleFunc("PATCH /teachers/", handlers.UpdateTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteMultipleTeachers)

	mux.HandleFunc("/students/", handlers.StudentHandler)

	mux.HandleFunc("/execs/", handlers.ExecHandler)

	return mux
}
