package services

import (
	"fmt"
	"strconv"
	"strings"
	"test-project/models"
	"test-project/utils"
)

type PersonService struct {
	people []models.Human
}

func NewPersonService() *PersonService {
	return &PersonService{
		people: []models.Human{},
	}
}

func (s *PersonService) CreatePerson() models.Person {
	name, ageInt, city, hobbies := s.readBasicInfo()
	person := models.Person{
		Name:    name,
		Age:     ageInt,
		City:    city,
		Hobbies: hobbies,
	}
	s.people = append(s.people, &person)
	return person
}

func (s *PersonService) CreateStudent() models.Student {
	name, ageInt, city, hobbies := s.readBasicInfo()
	course := utils.ReadInput("Digite o curso do aluno: ")
	semester := utils.ReadInput("Digite o semestre do aluno: ")
	semesterInt, _ := strconv.Atoi(semester)
	student := models.Student{
		Person: models.Person{
			Name:    name,
			Age:     ageInt,
			City:    city,
			Hobbies: hobbies,
		},
		Course:   course,
		Semester: semesterInt,
	}
	s.people = append(s.people, &student)
	return student
}

func (s *PersonService) CreateEmployee() models.Employee {
	name, ageInt, city, hobbies := s.readBasicInfo()
	salary := utils.ReadInput("Digite o salário do funcionário: ")
	salaryFloat, _ := strconv.ParseFloat(salary, 64)
	role := utils.ReadInput("Digite o cargo do funcionário: ")
	employee := models.Employee{
		Person: models.Person{
			Name:    name,
			Age:     ageInt,
			City:    city,
			Hobbies: hobbies,
		},
		Salary: salaryFloat,
		Role:   role,
	}
	s.people = append(s.people, employee)
	return employee
}

func (s *PersonService) ListPeople() {
	for i, human := range s.people {
		fmt.Printf("%d. %s\n", i+1, human.Describe())
	}
}

func (s *PersonService) readBasicInfo() (string, int, string, []string) {
	name := utils.ReadInput("Digite o nome da pessoa: ")

	var ageInt int

	for {
		age := utils.ReadInput("Digite a idade da pessoa: ")
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			fmt.Println("Idade inválida! Por favor, digite um número.")
			continue
		}
		if ageInt <= 0 || ageInt > 150 {
			fmt.Println("Idade deve estar entre 1 e 150 anos.")
			continue
		}
		break
	}

	city := utils.ReadInput("Digite a cidade da pessoa: ")
	hobbies := utils.ReadInput("Digite os hobbies da pessoa (separados por vírgula): ")
	return name, ageInt, city, strings.Split(hobbies, ",")
}
