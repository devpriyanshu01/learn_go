package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

var teachers = make(map[int]models.Teacher)

var nextID = 1

// initialize some dummy teachers data.
func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9F",
		Subject:   "Maths",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Karl",
		LastName:  "Marx",
		Class:     "9C",
		Subject:   "Physics",
	}
	nextID++
}

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeacherHandler(w, r)

	case http.MethodPost:
		addTeacherHandler(w, r)
	}
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
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

func getTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	trimmedValue := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id := strings.TrimSuffix(trimmedValue, "/")

	//get query params
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	//save fetched rows to a variable
	var teachersList []models.Teacher

	//get multiple teachers.
	if id == "" {
		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
		
		var args []interface{}
		if firstName != "" {
			query = query + " AND first_name = ?"
			args = append(args, firstName)
		}
		if lastName != "" {
			query = query + " AND last_name = ?"
			args = append(args, lastName)
		}

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
	} else {
		teacherID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("err", err)
			return
		}

		var teacher models.Teacher

		// err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		err = db.QueryRow("SELECT id, first_name, last_name, email FROM teachers WHERE id = ?", teacherID).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email)

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
		Count: len(teachersList),
	}
	//set content type
	w.Header().Set("Content-Type", "application/json")
	//encode data to json
	json.NewEncoder(w).Encode(response)
}
