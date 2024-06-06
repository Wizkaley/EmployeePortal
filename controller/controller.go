package controller

import (
	"EmployeeManagement/model"
	"EmployeeManagement/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		return
	}
	vars := mux.Vars(r)
	if vars["id"] != "" {
		employeeID, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}
		employee := repository.GetEmployee(employeeID)
		err = json.NewEncoder(w).Encode(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	page := r.URL.Query().Get("page")
	p, _ := strconv.Atoi(page)
	employees := repository.GetEmployeesRepo(p)
	json.NewEncoder(w).Encode(employees)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	var employee model.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repository.CreateEmployee(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	err = json.NewEncoder(w).Encode(employee.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		return
	}
	var employee model.Employee

	empID := mux.Vars(r)["id"]
	employeeID, err := strconv.Atoi(empID)
	if err != nil || empID == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = repository.UpdateEmployee(employeeID, employee); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		return
	}

	empID := mux.Vars(r)["id"]
	employeeID, err := strconv.Atoi(empID)
	if err != nil || empID == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.DeleteEmployee(employeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusGone)
}
