package commissionedClassification

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type SalesReceipt struct {
	// impl
}

type CommissionedClassification struct {
	commissionRate float64
	salary         float64
	salesReceipts  []SalesReceipt
}

func New(salary, commissionRate float64) CommissionedClassification {
	return CommissionedClassification{
		salary:         salary,
		commissionRate: commissionRate,
	}
}

func (c CommissionedClassification) CalculatePay(date date.Date) float64 {
	return 1.23 // impl
}

func (c CommissionedClassification) CommissionRate() float64 {
	return c.commissionRate
}

func (c CommissionedClassification) Salary() float64 {
	return c.salary
}
