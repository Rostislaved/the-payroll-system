package changeAddressStrategy

import "github.com/Rostislaved/the-payroll-system/employee"

type ChangeAddressStrategy struct {
	address string
}

func New(address string) ChangeAddressStrategy {
	return ChangeAddressStrategy{
		address: address,
	}
}

func (s ChangeAddressStrategy) Change(employeeIn employee.Employee) (employeeOut employee.Employee) {
	employeeIn.SetAddress(s.address)

	employeeOut = employeeIn

	return employeeOut
}
