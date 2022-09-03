package employee

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
)

type Employee struct {
	empID   int
	name    string
	address string

	classification PaymentClassification
	schedule       PaymentSchedule
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

type PaymentSchedule interface {
	IsPayday(date date.Date) bool
}

type PaymentClassification interface {
	CalculatePay(date date.Date) float64
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

func (e *Employee) Address() string {
	return e.address
}

func (e *Employee) Classification() PaymentClassification {
	return e.classification
}

func (e *Employee) SetClassification(classification PaymentClassification) {
	e.classification = classification
}

func (e *Employee) Schedule() PaymentSchedule {
	return e.schedule
}

func (e *Employee) SetSchedule(schedule PaymentSchedule) {
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
