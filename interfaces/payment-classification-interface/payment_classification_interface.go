package paymentClassificationInterface

import "github.com/Rostislaved/the-payroll-system/employee/date"

type PaymentClassification interface {
	CalculatePay(date date.Date) float64
}
