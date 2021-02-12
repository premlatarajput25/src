package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/premlatarajput25/StudentProject/controller"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getStudents", controller.GetStudents).Methods("GET")

	router.HandleFunc("/", controller.Starter).Methods("GET")
	router.HandleFunc("/addStudent", controller.PostStudent).Methods("POST")
	router.HandleFunc("/updateStudent", controller.PutStudent).Methods("PUT")
	router.HandleFunc("/deleteStudent/{id}", controller.DeleteStudent).Methods("DELETE")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
	// http.HandleFunc("/", starter)
	// http.HandleFunc("/getStudents", getStudents)
	// http.HandleFunc("/postStudent", postStudent)
	// http.HandleFunc("/putStudent", putStudent)
	// http.HandleFunc("/deleteStudent/{id}", deleteStudent)

}
