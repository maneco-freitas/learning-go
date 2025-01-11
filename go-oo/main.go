package main

import (
	"fmt"
	"test-project/services"
)

func main() {
	personService := services.NewPersonService()

	fmt.Println("Bem-vindo ao sistema de cadastro de pessoas!")
	person := personService.CreatePerson()
	fmt.Println("Pessoa cadastrada com sucesso!")
	fmt.Println(person.Describe())

	fmt.Println("Vamos cadastar um aluno!")
	student := personService.CreateStudent()
	fmt.Println("Aluno cadastrado com sucesso!")
	fmt.Println(student.Describe())

	fmt.Println("Vamos cadastar um funcionario!")
	employee := personService.CreateEmployee()
	fmt.Println("Funcionário cadastrado com sucesso!")
	fmt.Println(employee.Describe())

	fmt.Println("Lista de pessoas cadastradas:")
	personService.ListPeople()

	fmt.Println("Obrigado por usar o sistema de cadastro de pessoas!")
}
