package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
	"strings"
)

// check if any of the fields are empty.
func checkIfFieldsEmpty(body []byte) error {
	studentData := []models.Student{}
	err := json.Unmarshal(body, &studentData)
	fmt.Println("[]studentData", studentData)
	if err != nil {
		return utils.ErrorHandler(fmt.Errorf("ERROR UNMARSHALLING BODY TO MODELS.STUDENT"), "Error parsing the body")
	}
	for _, student := range studentData {
		//using reflect to iterate over all the fields.
		studentVal := reflect.ValueOf(student)

		for i := 0; i < studentVal.NumField(); i++ {
			field := studentVal.Field(i)
			if field.Interface() == "" {
				return fmt.Errorf("ONE OR MORE FIELDS ARE EMPTY")
			}
		}
	}
	return nil
}

// check if there are any unwanted fields.
func checkUnwantedFields(body []byte) error {
	var rawStudentData []map[string]interface{}

	err := json.Unmarshal(body, &rawStudentData)
	if err != nil {
		return utils.ErrorHandler(fmt.Errorf("ERROR UNMARSHALLING BODY TO RAWSTUDENTDATA"), "Error parsing the body")
	}

	//Get allowed fields
	var allowedFields []string

	studentType := reflect.TypeOf(models.Student{})
	for i := 0; i < studentType.NumField(); i++ {
		fieldTag := studentType.Field(i).Tag.Get("json")
		field := strings.TrimSuffix(fieldTag, ",omitempty")

		allowedFields = append(allowedFields, field)
	}

	//create a map to match to validate the fields.
	fieldValidator := map[string]struct{}{}
	for _, field := range allowedFields {
		fieldValidator[field] = struct{}{}
	}

	for _, student := range rawStudentData {
		for key := range student {
			_, ok := fieldValidator[key]
			if !ok {
				return utils.ErrorHandler(fmt.Errorf("UNWANTED FIELDS SENT"), "Unacceptable fields sent")
			}
		}
	}

	return nil
}

func CheckBlankFields(value interface{}) error {
	val := reflect.ValueOf(value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			return utils.ErrorHandler(errors.New("all fields are required"), "All fields are required")
		}
	}
	return nil
}

func GetFieldNames(model interface{}) []string {
	val := reflect.TypeOf(model)
	fields := []string{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldToAdd := strings.TrimSuffix(field.Tag.Get("json"), ",omitempty")
		fields = append(fields, fieldToAdd)
	}
	return fields
}
