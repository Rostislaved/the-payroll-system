package employee

import (
	"github.com/Rostislaved/the-payroll-system/employee/date"
	"github.com/Rostislaved/the-payroll-system/interfaces/affiliation-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-method-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
	"github.com/Rostislaved/the-payroll-system/paycheck"
)

type Employee struct {
	empID   int
	name    string
	address string

	classification paymentClassificationInterface.PaymentClassification
	schedule       paymentScheduleInterface.PaymentSchedule
	method         paymentMethodInterface.PaymentMethod
	affiliation    affiliationInterface.Affiliation
}

func (e *Employee) SetAffiliation(affiliation affiliationInterface.Affiliation) {
	e.affiliation = affiliation
}

func (e *Employee) Affiliation() (affiliation affiliationInterface.Affiliation) {
	return e.affiliation
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

func (e *Employee) Method() paymentMethodInterface.PaymentMethod {
	return e.method
}

func (e *Employee) SetMethod(method paymentMethodInterface.PaymentMethod) {
	e.method = method
}

func (e *Employee) IsPayDate(date date.Date) bool {
	return e.schedule.IsPayday(date)
}

func (e *Employee) Payday(pc paycheck.Paycheck) {
	gross := e.classification.CalculatePay(pc)
	deductions := e.affiliation.GetFee(pc)
	netPay := gross - deductions

	pc.SetGrossPay(gross)
	pc.SetDeductionsPay(deductions)
	pc.SetNetPay(netPay)

	e.Method().Pay(pc)
}

//func (e *Employee) CalculatePay(date date.Date) float64 {
//	pay := e.classification.CalculatePay(date)
//
//	// Affiliation GetFee(date)
//
//	return pay
//}
