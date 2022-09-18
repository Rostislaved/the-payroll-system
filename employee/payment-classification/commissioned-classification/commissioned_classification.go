package commissionedClassification

import (
	"github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification/sales-receipt"
	"github.com/Rostislaved/the-payroll-system/paycheck"
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

func (c *CommissionedClassification) CalculatePay(pc paycheck.Paycheck) float64 {
	panic("Not implemented")
	return 1.23
}

func (c *CommissionedClassification) CommissionRate() float64 {
	return c.commissionRate
}

func (c *CommissionedClassification) Salary() float64 {
	return c.salary
}
