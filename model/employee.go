package model

import "sync"

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

type Pagination struct {
	Next     int `json:"next_page"`
	Previous int `json:"previous_page"`
	Current  int `json:"current_page"`
	Limit    int `json:"limit"`
	Total    int `json:"total"`
}

type DB struct {
	MU        *sync.Mutex
	Employees []Employee
}
