package main

import (
	"StudentProject/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	//ctrl+c  to exit
}
