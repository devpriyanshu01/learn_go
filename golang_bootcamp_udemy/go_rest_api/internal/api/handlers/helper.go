package handlers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"restapi/internal/models"
	"restapi/pkg/utils"
	"strings"
)

// check if any of the fields are empty.
func checkIfFieldsEmpty(body []byte) error {
	fmt.Println("Checking if any of the fields are empty")
	studentData := models.Student{}
	err := json.Unmarshal(body, &studentData)
	if err != nil {
		return utils.ErrorHandler(fmt.Errorf("ERROR UNMARSHALLING BODY TO MODELS.STUDENT"), "Error parsing the body")
	}
	//using reflect to iterate over all the fields.
	studentVal := reflect.ValueOf(studentData)

	for i := 0; i < studentVal.NumField(); i++ {
		field := studentVal.Field(i)
		if field.Interface() == "" {
			return fmt.Errorf("ONE OR MORE FIELDS ARE EMPTY")
		}
	}
	return nil
}

// check if there are any unwanted fields.
func checkUnwantedFields(body []byte) error {
	var rawStudentData map[string]interface{}

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

	for key := range rawStudentData {
		_, ok := fieldValidator[key]
		if !ok {
			return utils.ErrorHandler(fmt.Errorf("UNWANTED FIELDS SENT"), "Unacceptable fields sent")
		}
	}

	return nil
}
