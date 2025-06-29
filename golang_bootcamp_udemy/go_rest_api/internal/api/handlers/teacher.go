package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb() //connect to db
	if err != nil {
		http.Error(w, "Error Connecting to Database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//store body in newTeacher struct
	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers) //json to struct
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		fmt.Println("error occured ===>>", err)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))

	//prepare the sql query
	stmt, err := db.Prepare("INSERT INTO teachers(first_name, last_name, email, class, subject) VALUES(?,?,?,?,?)")
	if err != nil {
		http.Error(w, "Error Preparing SQL Query", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	for i, teacher := range newTeachers {
		sqlResult, err := stmt.Exec(teacher.FirstName, teacher.LastName, teacher.Email, teacher.Class, teacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting data to database", http.StatusInternalServerError)
			return
		}
		insertedTeacherID, err := sqlResult.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last inserted ID", http.StatusInternalServerError)
			return
		}
		addedTeachers[i] = newTeachers[i]
		addedTeachers[i].ID = int(insertedTeacherID)

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
	//connect to database
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//try to get the teacherID
	trimmedValue := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id := strings.TrimSuffix(trimmedValue, "/")

	//save fetched rows to a variable
	var teachersList []models.Teacher

	//get multiple teachers as teacherID was not sent.
	if id == "" {
		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"

		var args []interface{}
		query, args = addFilters(query, r, args) //add filters functionality

		//add sorting functionality
		sortby := r.URL.Query().Get("sortby")
		splittedValue := strings.Split(sortby, ":")

		if len(splittedValue) >= 2 && validateSortingField(splittedValue) {
			query = query + " ORDER BY " + splittedValue[0] + " " + splittedValue[1]
		}

		//query the database
		sqlRows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "Error fetching teachers", http.StatusInternalServerError)
			fmt.Println("Error:----", err)
			return
		}
		defer sqlRows.Close()

		for sqlRows.Next() {
			var tchr models.Teacher
			sqlRows.Scan(&tchr.ID, &tchr.FirstName, &tchr.LastName, &tchr.Email, &tchr.Class, &tchr.Subject)
			teachersList = append(teachersList, tchr)
		}
	} else { //since teacherID was sent, find the teacher with given ID
		teacherID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("err", err)
			return
		}

		var teacher models.Teacher

		// err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

		if err == sql.ErrNoRows {
			http.Error(w, "teacher not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Error searching a teacher", http.StatusInternalServerError)
			fmt.Println("error:---", err)
			return
		}
		teachersList = append(teachersList, teacher)
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

func addFilters(query string, r *http.Request, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}
	//extract sent query params
	for key, value := range params {
		col := r.URL.Query().Get(key)
		if col != "" {
			query = query + " AND " + value + " = ?"
			args = append(args, col)
		}
	}
	return query, args
}

func validateSortingField(splittedValue []string) bool {
	//validate the sent sorting field.
	colFields := []string{"first_name", "last_name", "email", "class", "subject"}
	for _, field := range colFields {
		if splittedValue[0] == field {
			return true
		}
	}
	return false
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
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error Connecting to Database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherId).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher with given ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error fetching the database with given teacherID", http.StatusInternalServerError)
		return
	}

	//udpate the teacher in database
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", receivedTeacher.FirstName, receivedTeacher.LastName, receivedTeacher.Email, receivedTeacher.Class, receivedTeacher.Subject, existingTeacher.ID)
	if err != nil {
		http.Error(w, "Error updating teacher to database", http.StatusInternalServerError)
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

func UpdateTeacherFields(w http.ResponseWriter, r *http.Request) {
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

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(
		&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "No teacher found with ID sent", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Error fetching teacher from database", http.StatusInternalServerError)
		return
	}

	//update the existing teacher data with new values.
	/*
	for field, value := range toUpdate {
		switch field {
		case "first_name":
			existingTeacher.FirstName = value.(string)
		case "last_name":
			existingTeacher.LastName = value.(string)
		case "email":
			existingTeacher.Email = value.(string)
		case "class":
			existingTeacher.Class = value.(string)
		case "subject":
			existingTeacher.Subject = value.(string)	
		}
	}*/

	//code modification using reflect
	//for now consider only one field will be sent to update at once.
	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherVal.Type()
	for k, v := range toUpdate{
		fieldCount := teacherVal.NumField()
		for i := range fieldCount {
			field := teacherType.Field(i)
			if (k + ",omitempty") == field.Tag.Get("json") {
				teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))	
			}
		}
	}

	//update to database
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {
		http.Error(w, "Error updating data to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}

//delete teacher
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	reqUrl := r.URL.Path
	idStr := strings.TrimPrefix(reqUrl, "/teachers/")
	teacherID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlResult, err := db.Exec("DELETE FROM teachers WHERE id = ?", teacherID)
	if err != nil {
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return
	}

	affectedRow, err := sqlResult.RowsAffected()
	if affectedRow == 0 {
		http.Error(w, "No rows deleted", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error retrieving deleted result", http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		RowsAffected int64 `json:"rows_affected"`
	}{
		Status: "success",
		RowsAffected: affectedRow,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// w.WriteHeader(http.StatusNoContent) //on deletion we need to send it.
}


func UpdateTeachersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teachers multi - entry patch updater hit....")
	var toUpdate []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&toUpdate)
	if err != nil {
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}
	fmt.Println("PRINTING toUpdate", toUpdate)

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Error starting Transaction", http.StatusInternalServerError)
		fmt.Println("TRXN ERR", tx)
		return
	}

	for _, update := range toUpdate {
		fmt.Println("PRINTING update[id]", update["id"])
		fmt.Println("UPDATE ====>", update)
		id, ok := update["id"].(float64)
		if !ok {
			tx.Rollback()
			http.Error(w, "Invalid teacher id", http.StatusBadRequest)
			return
		}
		teacherID := int(id)
		fmt.Println("type of Id", reflect.TypeOf(teacherID))

		var oneTeacher models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(
			&oneTeacher.ID, &oneTeacher.FirstName, &oneTeacher.LastName, &oneTeacher.Email, &oneTeacher.Class, &oneTeacher.Subject)
		if err == sql.ErrNoRows {
			http.Error(w, "No Teacher Found", http.StatusInternalServerError)
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error fetching a teacher from database", http.StatusInternalServerError)
			return
		}

		fmt.Println("one teacher =======>", oneTeacher)
		teacherVal := reflect.ValueOf(&oneTeacher).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {	//skip updating the id field as id doesn't change
				continue
			}
			fieldCount := teacherVal.NumField()
			for i := range fieldCount {
				field := teacherType.Field(i)
				if (k + ",omitempty") == field.Tag.Get("json"){
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						value := reflect.ValueOf(v)
						if value.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(value.Convert(fieldVal.Type()))
						}else {
							tx.Rollback()
							log.Printf("can't convert %v to %v", value.Type(), fieldVal.Type())
							return
						}
					}else {
						log.Println("can't set the value", fieldVal)
						break
					}
				}else {
				}
			}
		}


		//update the db
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", oneTeacher.FirstName, 
		oneTeacher.LastName, oneTeacher.Email, oneTeacher.Class, oneTeacher.Subject, oneTeacher.ID)
		if err != nil {
			http.Error(w, "Error updating to database", http.StatusInternalServerError)
			tx.Rollback()
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)


}