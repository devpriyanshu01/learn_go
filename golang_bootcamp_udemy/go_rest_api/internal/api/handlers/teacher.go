package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {

	//Reject the request with fields not acceptable.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	var rawTeachers []map[string]interface{}
	err = json.Unmarshal(body, &rawTeachers)
	if err != nil {
		http.Error(w, "Error unmarshalling the body", http.StatusInternalServerError)
		fmt.Println("error -->", err)
		return
	}
	log.Println("rawTeachers:", rawTeachers)

	fields := []string{}
	teacherType := reflect.TypeOf(models.Teacher{})

	for i := 0; i < teacherType.NumField(); i++ {
		field := teacherType.Field(i).Tag.Get("json")
		field = strings.TrimSuffix(field, ",omitempty")

		fields = append(fields, field)
	}

	fmt.Println("fields", fields)

	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	for _, teacher := range rawTeachers {
		for key := range teacher {
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable fields found, send only allowed fields.", http.StatusBadRequest)
				return
			}
		}
	}

	//Check if none of the fields should be empty.
	var newTeachers []models.Teacher
	err = json.Unmarshal(body, &newTeachers)
	if err != nil {
		http.Error(w, "Error Unmarshalling the body to newTeachers", http.StatusInternalServerError)
		return
	}

	for _, teacher := range newTeachers {
		teacherVal := reflect.ValueOf(teacher)
		for i := 0; i < teacherVal.NumField(); i++ {
			fieldVal := teacherVal.Field(i)
			if fieldVal.Interface() == "" {
				http.Error(w, "No fields should be empty.", http.StatusBadRequest)
				return
			}
		}
	}

	addedTeachers, err := sqlconnect.AddTeachersDbHandler(newTeachers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	teachersList, err := sqlconnect.GetTeachersDbHandler(id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	err = sqlconnect.UpdateTeacherPutDbHandler(teacherId, receivedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	existingTeacher, err := sqlconnect.UpdateTeacherPatchDbHandler(teacherID, toUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	err = sqlconnect.UpdateMultipleTeachersDbHandler(toUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

func GetStudentsByTeachersId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	teacherID, err := strconv.Atoi(id)
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
		Status string           `json:"status,omitempty"`
		Count  int              `json:"count,omitempty"`
		Data   []models.Student `json:"data,omitempty"`
	}{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func GetStudentCountForaTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	teacherID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	studentCount, err := sqlconnect.GetStudentCountForaTeacherDbHandler(teacherID)
	if err != nil {
		http.Error(w, "failed to fetch the student count", http.StatusInternalServerError)
		return
	}

	type response struct {
		Status string
		Count int
	}
	res := response{
		Status: "success",
		Count: studentCount,
	}

	json.NewEncoder(w).Encode(res)
}
