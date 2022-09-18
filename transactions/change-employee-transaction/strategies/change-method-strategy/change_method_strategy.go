package changeMethodStrategy

import (
	"github.com/Rostislaved/the-payroll-system/employee"
	paymentMethodInterface "github.com/Rostislaved/the-payroll-system/interfaces/payment-method-interface"
)

type ChangeMethodStrategy struct {
	method paymentMethodInterface.PaymentMethod
}

func New(makeEmployeeStrategy paymentMethodInterface.PaymentMethod) ChangeMethodStrategy {
	return ChangeMethodStrategy{
		method: makeEmployeeStrategy,
	}
}

func (s ChangeMethodStrategy) Change(employeeIn employee.Employee) (employeeOut employee.Employee) {
	employeeOut = employeeIn

	method := s.method

	employeeOut.SetMethod(method)

	return employeeOut
}
