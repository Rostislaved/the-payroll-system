package changeNameStrategy

import "github.com/Rostislaved/the-payroll-system/employee"

type ChangeNameStrategy struct {
	name string
}

func New(name string) ChangeNameStrategy {
	return ChangeNameStrategy{
		name: name,
	}
}

func (s ChangeNameStrategy) Change(employeeIn employee.Employee) (employeeOut employee.Employee) {
	employeeIn.SetName(s.name)

	employeeOut = employeeIn

	return employeeOut
}
