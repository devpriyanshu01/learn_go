package sqlconnect

import (
	"encoding/json"
	"fmt"
	"restapi/internal/models"
	"restapi/pkg/utils"
)

func AddOneStudentDbHandler(studentData models.Student) (int, error) {
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

func AddStudentsDbHandler(body []byte) ([]int, error) {
	students := []models.Student{}
	err := json.Unmarshal(body, &students)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error parsing the body")
	}
	fmt.Println("students data:", students)

	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error establishing connection")
	}
	defer db.Close()

	//begin transaction
	txn, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error initiating txn.")
	}

	var insertedIds []int
	for _, student := range students {
		stmt, err := txn.Prepare("INSERT INTO students(first_name, last_name, email, class) VALUES(?, ?, ?, ?)")
		if err != nil {
			txn.Rollback()
			return nil, utils.ErrorHandler(err, "Error saving student")
		}
		defer stmt.Close()

		result, err := stmt.Exec(student.FirstName, student.LastName, student.Email, student.Class)
		if err != nil {
			txn.Rollback()
			return nil, utils.ErrorHandler(err, "Error saving student2")
		}

		insertedID, err := result.LastInsertId()
		if err != nil {
			txn.Rollback()
			return nil, utils.ErrorHandler(err, "Error saving student3")
		}
		insertedIds = append(insertedIds, int(insertedID))
	}
	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return nil, utils.ErrorHandler(err, "Error saving students4")
	}

	return insertedIds, nil
}

func DeleteOneStudentDbHandler(studentId int) (int, error) {
	db, err := ConnectDb()
	if err != nil {
		return 0, utils.ErrorHandler(err, "failed to delete student")
	}
	result, err := db.Exec("DELETE FROM students where id = ?", studentId)
	if err != nil {
		return 0, utils.ErrorHandler(err, "failed to delete student2")
	}
	noRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, utils.ErrorHandler(err, "failed to delete student3")
	}
	return int(noRowsAffected), nil
}

func GetOneStudentDbHandler(id int) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "failed to get student1")
	}

	sqlRow := db.QueryRow("SELECT first_name, last_name, email, class FROM students where id = ?", id)

	var student models.Student
	err = sqlRow.Scan(&student.FirstName, &student.LastName, &student.Email, &student.Class)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "error getting a student")
	}

	return student, nil

}

func GetStudentsDbHandler(ids []int) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return []models.Student{}, utils.ErrorHandler(err, "error getting students1")
	}

	fmt.Printf("printing received ids: %v\n", ids)

	var fetchedStudents []models.Student
	trx, err := db.Begin()
	if err != nil {
		return []models.Student{}, utils.ErrorHandler(err, "error getting students2")
	}

	for _, id := range ids {
		var student models.Student
		fmt.Println("processing for id:", id)
		sqlRow := db.QueryRow("SELECT first_name, last_name, email, class FROM students WHERE id = ?", id)
		err = sqlRow.Scan(&student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			trx.Rollback()
			return []models.Student{}, utils.ErrorHandler(err, "error getting students3")
		}
		fmt.Println("one student", student)
		fetchedStudents = append(fetchedStudents, student)
	}
	err = trx.Commit()
	if err != nil {
		return []models.Student{}, utils.ErrorHandler(err, "error getting students4")
	}
	return fetchedStudents, nil

}
