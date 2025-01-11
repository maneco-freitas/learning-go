package models

import "fmt"

type Student struct {
	Person
	Course   string
	Semester int
}

func (s Student) Describe() string {
	return fmt.Sprintf("%s está cursando %s no %dº semestre",
		s.Person.Describe(), s.Course, s.Semester)
}
