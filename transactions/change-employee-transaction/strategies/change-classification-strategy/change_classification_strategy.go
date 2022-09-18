package changeClassificationStrategy

import (
	"github.com/Rostislaved/the-payroll-system/employee"
	makeEmployeeStrategy "github.com/Rostislaved/the-payroll-system/interfaces/make-employee-strategy"
)

type ChangeClassificationStrategy struct {
	makeEmployeeStrategy makeEmployeeStrategy.MakeEmployeeStrategy
}

func New(makeEmployeeStrategy makeEmployeeStrategy.MakeEmployeeStrategy) ChangeClassificationStrategy {
	return ChangeClassificationStrategy{
		makeEmployeeStrategy: makeEmployeeStrategy,
	}
}

func (s ChangeClassificationStrategy) Change(employeeIn employee.Employee) (employeeOut employee.Employee) {
	classification := s.makeEmployeeStrategy.MakeClassification()
	schedule := s.makeEmployeeStrategy.MakeSchedule()

	employeeOut = employeeIn

	employeeOut.SetClassification(classification)
	employeeOut.SetSchedule(schedule)

	return employeeOut
}
