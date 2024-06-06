package controller

import (
	"EmployeeManagement/model"
	"EmployeeManagement/repository"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gorilla/mux"
)

func seedData() {
	repository.DataB = model.DB{
		MU: &sync.Mutex{},
		Employees: []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
			{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500},
		},
	}
}

func TestGetEmployeesHandler(t *testing.T) {
	seedData()
	tests := []struct {
		name    string
		url     string
		status  int
		expect  interface{}
		err     bool
		setVars map[string]string
	}{
		{
			name:   "Test GetEmployeesHandler",
			url:    "/employees?page=1",
			err:    false,
			status: http.StatusOK,
		},
		{
			name:    "Test GetEmployee",
			url:     "/employees",
			err:     false,
			status:  http.StatusOK,
			setVars: map[string]string{"id": "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			if tt.setVars != nil {
				r = mux.SetURLVars(r, tt.setVars)
			}
			w := httptest.NewRecorder()
			GetEmployees(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			_, _ = ioutil.ReadAll(resp.Body)
			if tt.err {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		})
	}
}

func TestCreateEmployeeHandler(t *testing.T) {
	seedData()
	tests := []struct {
		name   string
		url    string
		status int
		err    bool
	}{
		{
			name:   "Test CreateEmployeeHandler",
			url:    "/employees",
			err:    false,
			status: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			me, _ := json.Marshal(model.Employee{
				ID:       3,
				Name:     "John Doe",
				Position: "Senior Developer",
				Salary:   3500,
			})

			r, _ := http.NewRequest(http.MethodPost, tt.url, bytes.NewReader(me))
			w := httptest.NewRecorder()
			CreateEmployee(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			_, _ = ioutil.ReadAll(resp.Body)
			if tt.err {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	seedData()
	tests := []struct {
		name    string
		url     string
		status  int
		err     bool
		setVars map[string]string
	}{
		{
			name:    "Test UpdateEmployee",
			url:     "/employees/1",
			err:     false,
			status:  http.StatusOK,
			setVars: map[string]string{"id": "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			me, _ := json.Marshal(model.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Senior Developer",
				Salary:   3500,
			})

			r, _ := http.NewRequest(http.MethodPut, tt.url, bytes.NewReader(me))
			if tt.setVars != nil {
				r = mux.SetURLVars(r, tt.setVars)
			}
			w := httptest.NewRecorder()
			UpdateEmployee(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			_, _ = ioutil.ReadAll(resp.Body)
			if tt.err {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	seedData()
	tests := []struct {
		name    string
		url     string
		status  int
		err     bool
		setVars map[string]string
	}{
		{
			name:    "Test DeleteEmployee",
			url:     "/employees/1",
			err:     false,
			status:  http.StatusGone,
			setVars: map[string]string{"id": "1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, _ := http.NewRequest(http.MethodDelete, tt.url, nil)
			if tt.setVars != nil {
				r = mux.SetURLVars(r, tt.setVars)
			}
			w := httptest.NewRecorder()
			DeleteEmployee(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			_, _ = ioutil.ReadAll(resp.Body)
			if tt.err {
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusGone, resp.StatusCode)
			}
		})
	}
}
