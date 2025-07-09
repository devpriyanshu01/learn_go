package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
)

func AddOneStudent(w http.ResponseWriter, r *http.Request) {
	var studentData models.Student

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing the body", http.StatusInternalServerError)
		log.Println("Error is ----------", err)
		return
	}

	//Check if any of the fields are empty.
	err = checkIfFieldsEmpty(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check any unwanted fields sent.
	err = checkUnwantedFields(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &studentData)
	if err != nil {
		http.Error(w, "Error parsing the body", http.StatusBadRequest)
		return
	}
	studentId, err := sqlconnect.AddOneStudentDbHandler(studentData)
	if err != nil {
		http.Error(w, "Error adding a student.", http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("Student Created with Student_ID = %d", studentId)
	w.Write([]byte(response))
}

func AddStudents(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the sent data", http.StatusInternalServerError)
		return
	}

	err = checkIfFieldsEmpty(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = checkUnwantedFields(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertedStudentIds, err := sqlconnect.AddStudentsDbHandler(body)
	if err != nil {
		http.Error(w, "Error adding students, operation failed.", http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("Students added with IDs - %v", insertedStudentIds)
	w.Write([]byte(response))

}
