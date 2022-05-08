package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID		int
	Name		string
	Address		string
	DOB		time.Time
	Position	string
	Salary		int
	ManagaerID	int
}

func main() {
	var dilbert Employee
	var employeeOfTheMonth *Employee
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	employeeOfTheMonth = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (clever)"
	fmt.Printf("dilbert: %v\n", dilbert)
}
