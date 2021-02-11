package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/premlatarajput25/StudentProject/beans"
	configurationFile "github.com/premlatarajput25/StudentProject/configuration"
	"github.com/premlatarajput25/StudentProject/dataAccess"
)

func Starter(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, " Student Portal")

}

func GetStudents(response http.ResponseWriter, req *http.Request) { // []Student students
	fmt.Fprintf(response, "Available  Students Are \n")
	db, dberr := configurationFile.GetDBConnection()
	if dberr != nil {
		errorResponse(response, http.StatusInternalServerError, dberr.Error())
	} else {
		var studentAccess = dataAccess.StudentAccess{Db: db}
		students, fetchDataerr := studentAccess.GetAllStudents()

		if fetchDataerr != nil {
			errorResponse(response, http.StatusInternalServerError, fetchDataerr.Error())
		} else {
			jsonResponse(response, http.StatusOK, students)
		}
	}

}
func PutStudent(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Modifying  Students \n")
	var student beans.Student
	json.NewDecoder(req.Body).Decode(&student)
	db, dberr := configurationFile.GetDBConnection()
	if dberr != nil {
		errorResponse(response, http.StatusInternalServerError, dberr.Error())
	} else {
		var studentAccess = dataAccess.StudentAccess{Db: db}
		rowsAffected, fetchDataerr := studentAccess.UpdateStudent(&student)

		if fetchDataerr != nil {
			errorResponse(response, http.StatusInternalServerError, fetchDataerr.Error())
		} else if rowsAffected > 0 {
			data := fmt.Sprint(rowsAffected, " students were updated!!!")
			jsonResponse(response, http.StatusOK, data)
		} else {
			jsonResponse(response, http.StatusOK, "No students were updated!!!")
		}
	}
}
func PostStudent(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Adding  Students \n")
	var student beans.Student
	err := json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		errorResponse(response, http.StatusInternalServerError, err.Error())
	} else {
		db, dberr := configurationFile.GetDBConnection()
		if dberr != nil {
			errorResponse(response, http.StatusInternalServerError, dberr.Error())
		} else {
			//

			var studentAccess = dataAccess.StudentAccess{Db: db}
			fetchDataerr := studentAccess.InsertStudent(&student)

			if fetchDataerr != nil {
				errorResponse(response, http.StatusInternalServerError, fetchDataerr.Error())
			} else {
				fmt.Println(student)
				jsonResponse(response, http.StatusOK, student)
			}
		}
	}

}
func DeleteStudent(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "Removing  Students \n")
	values := mux.Vars(req)
	ids := values["id"]
	fmt.Println(ids)
	id, err := strconv.ParseInt(ids, 10, 64)
	// json.NewDecoder(req.)
	if err != nil {
		errorResponse(response, http.StatusInternalServerError, err.Error())
	} else {
		db, dberr := configurationFile.GetDBConnection()
		if dberr != nil {
			errorResponse(response, http.StatusInternalServerError, dberr.Error())
		} else {
			var studentAccess = dataAccess.StudentAccess{Db: db}
			students, fetchDataerr := studentAccess.DeleteStudent(id)

			if fetchDataerr != nil {
				errorResponse(response, http.StatusInternalServerError, fetchDataerr.Error())
			} else {
				jsonResponse(response, http.StatusOK, students)
			}
		}
	}
}

func errorResponse(response http.ResponseWriter, errorCode int, errorMsg string) {
	jsonResponse(response, http.StatusInternalServerError, map[string]string{"error": errorMsg})
}
func jsonResponse(response http.ResponseWriter, code int, message interface{}) {
	data, _ := json.Marshal(message)
	response.Write(data)
	response.WriteHeader(code)
	response.Header().Set("Content-Type", "application/json")
}
