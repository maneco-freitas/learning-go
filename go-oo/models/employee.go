package models

import "fmt"

type Employee struct {
	Person
	Salary float64
	Role   string
}

func (e Employee) Describe() string {
	return fmt.Sprintf("%s trabalha com %s com salário de R$%.2f",
		e.Person.Describe(), e.Role, e.Salary)
}
