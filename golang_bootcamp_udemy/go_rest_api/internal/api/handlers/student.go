package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
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

func DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Student ID, it should be a a number", http.StatusBadRequest)
		return
	}
	row, err := sqlconnect.DeleteOneStudentDbHandler(id)
	if err != nil {
		http.Error(w, "Error deleting student", http.StatusInternalServerError)
		return
	}
	response := fmt.Sprintf("Deleted Successfully %d row", row)
	w.Write([]byte(response))
}

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
	}

	student, err := sqlconnect.GetOneStudentDbHandler(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&student)

}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	studentIds := []int{}
	err := json.NewDecoder(r.Body).Decode(&studentIds)
	if err != nil {
		http.Error(w, "error parsing the body", http.StatusBadRequest)
		return
	}

	students, err := sqlconnect.GetStudentsDbHandler(studentIds)
	fmt.Println("final students", students)
	if err != nil {
		http.Error(w, "error fetching studetns", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "error encoding to response", http.StatusInternalServerError)
		return
	}

}

func UpdateOneStudent(w http.ResponseWriter, r *http.Request) {
	var toUpdate map[string]interface{}
	fmt.Println("inside update one student")
	err := json.NewDecoder(r.Body).Decode(&toUpdate)
	fmt.Println("toUpdate", toUpdate)
	if err != nil {
		http.Error(w, "error parsing the body", http.StatusBadRequest)
		return
	}

	err = sqlconnect.UpdateOneStudentDbHandler(toUpdate)
	if err != nil {
		http.Error(w, "failed to update student", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("student updated"))
}

func GetStudentsByTeachersId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	teacherID, err  := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}
	var students []models.Student
	students, err = sqlconnect.GetStudentsByTeacherIdDbHandler(teacherID)
	if err != nil {
		http.Error(w, "failed to get students list", http.StatusInternalServerError)
		return
	}
	fmt.Println("printing students:", students)
	fmt.Println("len of students", len(students))
	
	response := struct {
		Status string `json:"status,omitempty"`
		Count int `json:"count,omitempty"`
		Data []models.Student `json:"data,omitempty"`
	}{
		Status: "success",
		Count: len(students),
		Data: students,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}


