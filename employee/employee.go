package employee

import (
	"time"
)

type Employee struct {
	empid   int
	name    string
	address string

	classification PaymentClassification
	schedule       PaymentSchedule
	method         PaymentMethod
	affiliation    Affiliation
}

type Affiliation interface {
	// impl
}

type PaymentMethod interface {
	Pay(amount int)
}

type PaymentSchedule interface {
	IsPayday(date time.Time) bool
}

type PaymentClassification interface {
	CalculatePay(date time.Time) float32
}

func New(empid int, name, address string) Employee {
	return Employee{
		empid:   empid,
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

func (e *Employee) IsPayDate(date time.Time) bool {
	return e.schedule.IsPayday(date)
}

func (e *Employee) CalculatePay(date time.Time) float32 {
	pay := e.classification.CalculatePay(date)

	// Affilation GetFee(date)

	return pay
}

func (e *Employee) Payday() {
	panic("implement me")
}
