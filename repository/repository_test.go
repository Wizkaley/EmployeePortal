package repository

import (
	"EmployeeManagement/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var DataB = model.DB{
// 	MU: &sync.Mutex{},
// 	Employees: []model.Employee{
// 		{
// 			ID:       1,
// 			Name:     "John Doe",
// 			Position: "Senior Developer",
// 			Salary:   3500,
// 		},
// 		{
// 			ID:       2,
// 			Name:     "Jane Doe",
// 			Position: "Junior Developer",
// 			Salary:   2500,
// 		},
// 	},
// }

// func SetupTestData()	{
// 	DataB :=
// }

// func init() {
// 	SetupTestData()
// }

func TestGetEmployee(t *testing.T) {

	t.Run("GetEmployee", func(t *testing.T) {
		employee := GetEmployee(1)
		if employee.ID != 1 {
			t.Errorf("Expected ID to be 1, but got %d", employee.ID)
		}

		assert.Equal(t, 1, employee.ID)

	})

	t.Run("GetEmployeeNotFound", func(t *testing.T) {
		employee := GetEmployee(10)

		assert.Empty(t, employee)

	})

	t.Run("GetEmployees", func(t *testing.T) {
		DataB.MU = &sync.Mutex{}
		DataB.Employees = []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
			{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500},
		}
		employees := GetEmployeesRepo(1)
		emps, ok := employees.(map[string]interface{})
		if !ok {
			t.Errorf("Expected employees to be of type []model.Employee, but got %T", ok)
		}
		e, OK := emps["employees"].([]model.Employee)
		if !OK {
			t.Errorf("Expected employees to be of type []model.Employee, but got %T", OK)
		}
		if len(e) != 2 {
			t.Errorf("Expected len to be 2, but got %v", e)
		}

		assert.Len(t, e, 2)
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("DeleteEmployee", func(t *testing.T) {
		// DataB.MU = &sync.Mutex{}
		// DataB.Employees = []model.Employee{
		// 	{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
		// 	{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500},
		// }
		err := DeleteEmployee(1)
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}

		assert.Contains(t, DataB.Employees, model.Employee{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500})
		assert.NotContains(t, DataB.Employees, model.Employee{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500})
		assert.Nil(t, err)

	})

	t.Run("DeleteEmployeeNotFound", func(t *testing.T) {
		DataB.MU = &sync.Mutex{}
		DataB.Employees = []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
		}
		err := DeleteEmployee(10)

		// we dont have any errors from db at this point only thing we can check is if the employee is not found in the slice
		assert.Nil(t, err)

	})
}

func TestUpdateEmployee(t *testing.T) {
	t.Run("UpdateEmployee", func(t *testing.T) {
		DataB.MU = &sync.Mutex{}
		DataB.Employees = []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
			{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500},
		}
		err := UpdateEmployee(1, model.Employee{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 4500})
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}
		assert.Equal(t, float64(4500), DataB.Employees[0].Salary)
		assert.Nil(t, err)
	})

	t.Run("UpdateEmployeeNotFound", func(t *testing.T) {
		DataB.MU = &sync.Mutex{}
		DataB.Employees = []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
		}
		err := UpdateEmployee(10, model.Employee{ID: 10, Name: "John Doe", Position: "Senior Developer", Salary: 4500})
		// we dont have any errors from db at this point only thing we can check is if the employee is not found in the slice
		assert.Nil(t, err)
	})
}

func TestCreateEmployee(t *testing.T) {
	t.Run("CreateEmployee", func(t *testing.T) {
		DataB.MU = &sync.Mutex{}
		DataB.Employees = []model.Employee{
			{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
		}
		err := CreateEmployee(model.Employee{ID: 4, Name: "Pane Doe", Position: "Junior Developer", Salary: 2500})
		if err != nil {
			t.Errorf("Expected err to be nil, but got %v", err)
		}
		assert.Contains(t, DataB.Employees, model.Employee{ID: 4, Name: "Pane Doe", Position: "Junior Developer", Salary: 2500})
		assert.Nil(t, err)
	})
}
