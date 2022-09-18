package changeAffiliationStrategy

import (
	"github.com/Rostislaved/the-payroll-system/employee"
	affiliationInterface "github.com/Rostislaved/the-payroll-system/interfaces/affiliation-interface"
)

type ChangeAffiliationStrategy struct {
	affiliation affiliationInterface.Affiliation
}

func New(affiliation affiliationInterface.Affiliation) ChangeAffiliationStrategy {
	return ChangeAffiliationStrategy{
		affiliation: affiliation,
	}
}

func (s ChangeAffiliationStrategy) Change(employeeIn employee.Employee) (employeeOut employee.Employee) {
	employeeOut = employeeIn

	employeeOut.SetAffiliation(s.affiliation)

	return employeeOut
}
