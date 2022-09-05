package employee

import (
	"github.com/Rostislaved/the-payroll-system/employee/date"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
)

type Employee struct {
	empID   int
	name    string
	address string

	classification paymentClassificationInterface.PaymentClassification
	schedule       paymentScheduleInterface.PaymentSchedule
	method         PaymentMethod
	affiliation    Affiliation
}

func (e *Employee) SetAffiliation(affiliation Affiliation) {
	e.affiliation = affiliation
}

func (e *Employee) Affiliation() (affiliation Affiliation) {
	return e.affiliation
}

type Affiliation interface {
	GetFee(date date.Date) (fee float64)
}

type PaymentMethod interface {
	Pay(amount int)
}

func New(empID int, name, address string) Employee {
	return Employee{
		empID:   empID,
		name:    name,
		address: address,
	}
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) Address() string {
	return e.address
}

func (e *Employee) SetAddress(address string) {
	e.address = address
}

func (e *Employee) Classification() paymentClassificationInterface.PaymentClassification {
	return e.classification
}

func (e *Employee) SetClassification(classification paymentClassificationInterface.PaymentClassification) {
	e.classification = classification
}

func (e *Employee) Schedule() paymentScheduleInterface.PaymentSchedule {
	return e.schedule
}

func (e *Employee) SetSchedule(schedule paymentScheduleInterface.PaymentSchedule) {
	e.schedule = schedule
}

func (e *Employee) Method() PaymentMethod {
	return e.method
}

func (e *Employee) SetMethod(method PaymentMethod) {
	e.method = method
}

func (e *Employee) IsPayDate(date date.Date) bool {
	return e.schedule.IsPayday(date)
}

func (e *Employee) CalculatePay(date date.Date) float64 {
	pay := e.classification.CalculatePay(date)

	// Affilation GetFee(date)

	return pay
}

func (e *Employee) Payday() {
	panic("implement me")
}
