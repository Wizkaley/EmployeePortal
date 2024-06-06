package main

import (
	"EmployeeManagement/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	router := mux.NewRouter()

	router.HandleFunc("/employees/{id}", controller.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", controller.DeleteEmployee).Methods("DELETE")
	// http.HandleFunc("/employees", controller.GetEmployees)
	// http.Handle("/employees/{id}", controller.GetEmployee).Methods("GET")
	// http.HandleFunc("/employees/{id}", controller.CreateEmployee)
	// http.HandleFunc("/employees/{id}", controller.UpdateEmployee)
	// http.HandleFunc("/employees/{id}", controller.DeleteEmployee)
	// http.ListenAndServe(":8080", nil)

	log.Fatal(http.ListenAndServe(":8080", router))
}
