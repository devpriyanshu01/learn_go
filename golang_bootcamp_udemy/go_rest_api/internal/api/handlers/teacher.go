package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	//store body in newTeacher struct
	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers) //json to struct
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		fmt.Println("error occured ===>>", err)
		return
	}

	addedTeachers, shouldReturn := sqlconnect.AddTeachersDbHandler(w, newTeachers)
	if shouldReturn {
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(newTeachers),
		Data:   addedTeachers,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {
	//try to get the teacherID
	trimmedValue := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id := strings.TrimSuffix(trimmedValue, "/")

	//connect to database
	teachersList, shouldReturn := sqlconnect.GetTeachersDbHandler(w, id, r)
	if shouldReturn {
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Data:   teachersList,
		Count:  len(teachersList),
	}
	//set content type
	w.Header().Set("Content-Type", "application/json")
	//encode data to json
	json.NewEncoder(w).Encode(response)
}

// PUT - means to update complete row
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	reqUrl := r.URL.Path
	idStr := strings.TrimPrefix(reqUrl, "/teachers/")
	fmt.Println("id", idStr)
	teacherId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var receivedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&receivedTeacher)
	if err != nil {
		http.Error(w, "Error Decoding Sent Body", http.StatusBadRequest)
		return
	}

	//connect to Database
	shouldReturn := sqlconnect.UpdateTeacherPutDbHandler(w, teacherId, receivedTeacher)
	if shouldReturn {
		return
	}

	response := struct {
		Status string `json:"status"`
	}{
		Status: "success, Teacher was udpdated",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateTeacherFieldsPatch(w http.ResponseWriter, r *http.Request) {
	//get the sent id in url
	id := r.PathValue("id")
	// idStr := strings.TrimPrefix(reqUrl, "/teachers/")
	teacherID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error Parsing the sent teacher ID", http.StatusBadRequest)
		return
	}

	var toUpdate map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&toUpdate)
	if err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}
	fmt.Println("to udpate", toUpdate)

	existingTeacher, shouldReturn := sqlconnect.UpdateTeacherPatchDbHandler(w, teacherID, toUpdate)
	if shouldReturn {
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}

// delete teacher
func DeleteOneTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	teacherID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	affectedRow, shouldReturn := sqlconnect.DeleteOneTeacherDbHandler(w, teacherID)
	if shouldReturn {
		return
	}

	response := struct {
		Status       string `json:"status"`
		RowsAffected int64  `json:"rows_affected"`
	}{
		Status:       "success",
		RowsAffected: affectedRow,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// w.WriteHeader(http.StatusNoContent) //on deletion we need to send it.
}

func UpdateTeachersHandler(w http.ResponseWriter, r *http.Request) {
	var toUpdate []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&toUpdate)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}
	fmt.Println("PRINTING toUpdate", toUpdate)

	shouldReturn := sqlconnect.UpdateMultipleTeachersDbHandler(w, toUpdate)
	if shouldReturn {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteMultipleTeachers(w http.ResponseWriter, r *http.Request) {
	var idsToDelete []int
	json.NewDecoder(r.Body).Decode(&idsToDelete)
	if len(idsToDelete) < 1 {
		http.Error(w, "No IDs were sent for deletion", http.StatusBadRequest)
		return
	}

	deletedIds, shouldReturn := sqlconnect.DeleteMultipleTeachersDbHandler(w, idsToDelete)
	if shouldReturn {
		return
	}

	response := struct {
		Status     string `json:"status,omitempty"`
		DeletedIds []int  `json:"deletedIds,omitempty"`
	}{
		Status:     "Successfully Deleted Teachers",
		DeletedIds: deletedIds,
	}
	json.NewEncoder(w).Encode(response)
}
