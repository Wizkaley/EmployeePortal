package repository

import (
	"EmployeeManagement/model"
	"sync"
)

var DataB model.DB

var employees = []model.Employee{
	{ID: 1, Name: "John Doe", Position: "Senior Developer", Salary: 3500},
	{ID: 2, Name: "Jane Doe", Position: "Junior Developer", Salary: 2500},
	{ID: 3, Name: "Pane Doe", Position: "Junior Developer", Salary: 2500},
	{ID: 4, Name: "Lane Doe", Position: "Junior Developer", Salary: 2500},
}

func init() {
	DataB = model.DB{
		MU:        &sync.Mutex{},
		Employees: employees}
}

func GetEmployee(id int) model.Employee {
	DataB.MU.Lock()
	for _, e := range DataB.Employees {
		if e.ID == id {
			v := e
			DataB.MU.Unlock()
			return v
		}
	}
	DataB.MU.Unlock()

	return model.Employee{}
}

func GetEmployeesRepo(page int) interface{} {

	mp := make(map[string]interface{})

	employees, pageData := Paginate(page, 2, len(DataB.Employees))
	mp["employees"] = employees
	mp["page"] = pageData

	return mp
}

func CreateEmployee(employee model.Employee) error {

	DataB.MU.Lock()
	DataB.Employees = append(DataB.Employees, employee)
	DataB.MU.Unlock()
	// any db errors at this point would be caught here
	return nil
}

func UpdateEmployee(employeeID int, employee model.Employee) error {

	for i, e := range DataB.Employees {
		if e.ID == employeeID {
			DataB.MU.Lock()
			DataB.Employees[i] = employee
			DataB.MU.Unlock()
		}
	}

	// any db errors at this point would be caught here
	return nil
}

func DeleteEmployee(id int) error {
	for i, v := range DataB.Employees {
		if v.ID == id {
			DataB.MU.Lock()
			DataB.Employees = append(DataB.Employees[:i], DataB.Employees[i+1:]...)
			DataB.MU.Unlock()
		}
	}

	// any db errors at this point would be caught here
	return nil
}

func Paginate(pageNum int, pageSize int, empLen int) (employees []model.Employee, pageData *model.Pagination) {
	if pageNum >= empLen {
		return
	}
	start := (pageNum - 1) * pageSize
	end := pageNum * pageSize
	if end > empLen {
		end = empLen
	}

	pageData = &model.Pagination{}
	total := len(employees) / pageSize
	// total := len(employees)

	rem := len(employees) % pageSize
	if rem == 0 {
		pageData.Total = total
	} else {
		pageData.Total = total + 1
	}

	pageData.Current = pageNum
	pageData.Limit = pageSize

	if pageNum < 0 {
		pageData.Next = pageNum + 1
	} else if pageNum <= pageData.Total {
		pageData.Previous = pageNum - 1
		pageData.Next = pageNum + 1
	} else if pageNum == pageData.Total {
		pageData.Next = 0
		pageData.Previous = pageNum - 1
	}

	DataB.MU.Lock()
	employees = DataB.Employees[start:end]
	DataB.MU.Unlock()
	return
}
