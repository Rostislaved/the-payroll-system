package commissionedClassification

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/commissioned-classification/sales-receipt"
)

type CommissionedClassification struct {
	commissionRate float64
	salary         float64
	salesReceipts  []salesReceipt.SalesReceipt
}

func (c *CommissionedClassification) SalesReceipts() []salesReceipt.SalesReceipt {
	return c.salesReceipts
}

func New(salary, commissionRate float64) *CommissionedClassification {
	return &CommissionedClassification{
		salary:         salary,
		commissionRate: commissionRate,
	}
}

func (c *CommissionedClassification) AddSalesReceipt(salesReceipt salesReceipt.SalesReceipt) {
	c.salesReceipts = append(c.salesReceipts, salesReceipt)
}

func (c *CommissionedClassification) CalculatePay(date date.Date) float64 {
	return 1.23 // impl
}

func (c *CommissionedClassification) CommissionRate() float64 {
	return c.commissionRate
}

func (c *CommissionedClassification) Salary() float64 {
	return c.salary
}
