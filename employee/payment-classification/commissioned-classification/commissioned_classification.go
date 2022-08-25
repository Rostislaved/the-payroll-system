package commissionedClassification

import "time"

type SalesReceipt struct {
	// impl
}

type CommissionedClassification struct {
	commissionRate float32
	salary         float32
	salesReceipts  []SalesReceipt
}

func New(salary, commissionRate float32) CommissionedClassification {
	return CommissionedClassification{
		salary:         salary,
		commissionRate: commissionRate,
	}
}

func (c CommissionedClassification) CalculatePay(date time.Time) float32 {
	return 1.23 // impl
}

func (c CommissionedClassification) CommissionRate() float32 {
	return c.commissionRate
}

func (c CommissionedClassification) Salary() float32 {
	return c.salary
}
