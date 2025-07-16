package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func teacherRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /teachers/", handlers.GetTeacherHandler)
	mux.HandleFunc("POST /teachers/", handlers.AddTeacherHandler)
	mux.HandleFunc("PUT /teachers/", handlers.UpdateTeacher)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.UpdateTeacherFieldsPatch)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacher)
	mux.HandleFunc("PATCH /teachers/", handlers.UpdateTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteMultipleTeachers)
	mux.HandleFunc("GET /teachers/{id}/students", handlers.GetStudentsByTeachersId)
	mux.HandleFunc("GET /teachers/{id}/studentcount", handlers.GetStudentCountForaTeacher)

	return mux
}
