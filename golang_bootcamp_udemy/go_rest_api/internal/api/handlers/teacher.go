package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
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
		getTeachersHandler(w, r)

	case http.MethodPost:
		addTeacherHandler(w, r)
	}
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()
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

	// var addedTeachers []models.Teacher
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

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// trimmedValue := strings.TrimPrefix(r.URL.Path, "/teachers/")
	// id := strings.TrimSuffix(trimmedValue, "/")
	// teacherID, err := strconv.Atoi(id)
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }

	teachersList := make([]models.Teacher, 0, len(teachers))

	for _, teacher := range teachers {
		teachersList = append(teachersList, teacher)
		// if teacher.ID == teacherID {
		// 	teachersList = append(teachersList, teacher)
		// }
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachersList),
		Data:   teachersList,
	}
	//set content type
	w.Header().Set("Content-Type", "application/json")
	//encode data to json
	json.NewEncoder(w).Encode(response)
}
