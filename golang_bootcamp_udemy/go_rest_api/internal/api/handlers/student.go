package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
)

func AddOneStudent(w http.ResponseWriter, r *http.Request) {
	var studentData models.Student

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error parsing the body", http.StatusInternalServerError)
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
