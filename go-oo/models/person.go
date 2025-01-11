package models

import (
	"fmt"
	"strings"
)

type Person struct {
	Name    string
	Age     int
	City    string
	Hobbies []string
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s tem %d anos, mora em %s e gosta de %s",
		p.Name, p.Age, p.City, strings.Join(p.Hobbies, ", "))
}
