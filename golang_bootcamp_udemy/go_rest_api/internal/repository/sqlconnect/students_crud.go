package sqlconnect

import (
	"restapi/internal/models"
	"restapi/pkg/utils"
)

func AddOneStudentDbHandler(studentData models.Student) (int, error){
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Connection to server failed.")
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO students(first_name, last_name, email, class) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error saving student.")
	}
	defer stmt.Close()

	result, err := stmt.Exec(studentData.FirstName, studentData.LastName, studentData.Email, studentData.Class)
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error saving student.")
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		return 0, utils.ErrorHandler(err, "Error saving student.")
	}
	
	return int(insertedId), nil
}