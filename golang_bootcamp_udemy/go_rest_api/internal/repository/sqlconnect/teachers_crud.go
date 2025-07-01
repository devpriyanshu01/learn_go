package sqlconnect

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"strconv"
	"strings"
)

func AddTeachersDbHandler(w http.ResponseWriter, newTeachers []models.Teacher) ([]models.Teacher, bool) {
	db, err := ConnectDb() //connect to db
	if err != nil {
		http.Error(w, "Error Connecting to Database", http.StatusInternalServerError)
		return nil, true
	}
	defer db.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))

	//prepare the sql query
	stmt, err := db.Prepare("INSERT INTO teachers(first_name, last_name, email, class, subject) VALUES(?,?,?,?,?)")
	if err != nil {
		http.Error(w, "Error Preparing SQL Query", http.StatusInternalServerError)
		return nil, true
	}
	defer stmt.Close()

	for i, teacher := range newTeachers {
		sqlResult, err := stmt.Exec(teacher.FirstName, teacher.LastName, teacher.Email, teacher.Class, teacher.Subject)
		if err != nil {
			http.Error(w, "Error inserting data to database", http.StatusInternalServerError)
			return nil, true
		}
		insertedTeacherID, err := sqlResult.LastInsertId()
		if err != nil {
			http.Error(w, "Error getting last inserted ID", http.StatusInternalServerError)
			return nil, true
		}
		addedTeachers[i] = newTeachers[i]
		addedTeachers[i].ID = int(insertedTeacherID)

	}
	return addedTeachers, false
}

func GetTeachersDbHandler(w http.ResponseWriter, id string, r *http.Request) ([]models.Teacher, bool) {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, true
	}
	defer db.Close()

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
			return nil, true
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
			return nil, true
		}

		var teacher models.Teacher

		// err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

		if err == sql.ErrNoRows {
			http.Error(w, "teacher not found", http.StatusNotFound)
			return nil, true
		} else if err != nil {
			http.Error(w, "Error searching a teacher", http.StatusInternalServerError)
			fmt.Println("error:---", err)
			return nil, true
		}
		teachersList = append(teachersList, teacher)
	}
	return teachersList, false
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

func UpdateTeacherPutDbHandler(w http.ResponseWriter, teacherId int, receivedTeacher models.Teacher) bool {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error Connecting to Database", http.StatusInternalServerError)
		return true
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherId).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher with given ID not found", http.StatusNotFound)
		return true
	}
	if err != nil {
		http.Error(w, "Error fetching the database with given teacherID", http.StatusInternalServerError)
		return true
	}

	//udpate the teacher in database
	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", receivedTeacher.FirstName, receivedTeacher.LastName, receivedTeacher.Email, receivedTeacher.Class, receivedTeacher.Subject, existingTeacher.ID)
	if err != nil {
		http.Error(w, "Error updating teacher to database", http.StatusInternalServerError)
		return true
	}
	return false
}

func UpdateTeacherPatchDbHandler(w http.ResponseWriter, teacherID int, toUpdate map[string]interface{}) (models.Teacher, bool) {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return models.Teacher{}, true
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(
		&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "No teacher found with ID sent", http.StatusBadRequest)
		return models.Teacher{}, true
	}
	if err != nil {
		http.Error(w, "Error fetching teacher from database", http.StatusInternalServerError)
		return models.

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
				}*/Teacher{}, true
	}

	//code modification using reflect
	//for now consider only one field will be sent to update at once.
	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	teacherType := teacherVal.Type()
	for k, v := range toUpdate {
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
		return models.Teacher{}, true
	}
	return existingTeacher, false
}

func DeleteOneTeacherDbHandler(w http.ResponseWriter, teacherID int) (int64, bool) {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return 0, true
	}
	defer db.Close()

	sqlResult, err := db.Exec("DELETE FROM teachers WHERE id = ?", teacherID)
	if err != nil {
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return 0, true
	}

	affectedRow, err := sqlResult.RowsAffected()
	if affectedRow == 0 {
		http.Error(w, "No rows deleted", http.StatusNotFound)
		return 0, true
	}
	if err != nil {
		http.Error(w, "Error retrieving deleted result", http.StatusInternalServerError)
		return 0, true
	}
	return affectedRow, false
}

func UpdateMultipleTeachersDbHandler(w http.ResponseWriter, toUpdate []map[string]interface{}) bool {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return true
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Error starting Transaction", http.StatusInternalServerError)
		fmt.Println("TRXN ERR", tx)
		return true
	}

	for _, update := range toUpdate {
		fmt.Println("PRINTING update[id]", update["id"])
		fmt.Println("UPDATE ====>", update)
		id, ok := update["id"].(float64)
		if !ok {
			tx.Rollback()
			http.Error(w, "Invalid teacher id", http.StatusBadRequest)
			return true
		}
		teacherID := int(id)
		fmt.Println("type of Id", reflect.TypeOf(teacherID))

		var oneTeacher models.Teacher
		err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(
			&oneTeacher.ID, &oneTeacher.FirstName, &oneTeacher.LastName, &oneTeacher.Email, &oneTeacher.Class, &oneTeacher.Subject)
		if err == sql.ErrNoRows {
			http.Error(w, "No Teacher Found", http.StatusInternalServerError)
			tx.Rollback()
			return true
		}
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error fetching a teacher from database", http.StatusInternalServerError)
			return true
		}

		fmt.Println("one teacher =======>", oneTeacher)
		teacherVal := reflect.ValueOf(&oneTeacher).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" { //skip updating the id field as id doesn't change
				continue
			}
			fieldCount := teacherVal.NumField()
			for i := range fieldCount {
				field := teacherType.Field(i)
				if (k + ",omitempty") == field.Tag.Get("json") {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						value := reflect.ValueOf(v)
						if value.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(value.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("can't convert %v to %v", value.Type(), fieldVal.Type())
							return true
						}
					} else {
						log.Println("can't set the value", fieldVal)
						break
					}
				} else {
				}
			}
		}

		//update the db
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", oneTeacher.FirstName,
			oneTeacher.LastName, oneTeacher.Email, oneTeacher.Class, oneTeacher.Subject, oneTeacher.ID)
		if err != nil {
			http.Error(w, "Error updating to database", http.StatusInternalServerError)
			tx.Rollback()
			return true
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return true
	}
	return false
}

func DeleteMultipleTeachersDbHandler(w http.ResponseWriter, idsToDelete []int) ([]int, bool) {
	db, err := ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return nil, true
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, "Error initiating transaction", http.StatusInternalServerError)
		return nil, true
	}

	deletedIds := make([]int, len(idsToDelete))
	for i, id := range idsToDelete {

		result, err := tx.Exec("DELETE FROM teachers WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
			return nil, true
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, "Error getting rows affected", http.StatusInternalServerError)
			return nil, true
		}
		if rowsAffected > 0 {
			deletedIds[i] = id
		}
		if rowsAffected < 1 {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("teacher with ID = %d not found/deleted", id), http.StatusInternalServerError)
			return nil, true
		}
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, "Error committing transactions", http.StatusInternalServerError)
		return nil, true
	}
	return deletedIds, false
}
