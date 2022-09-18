package paymentClassificationInterface

import (
	"github.com/Rostislaved/the-payroll-system/paycheck"
)

type PaymentClassification interface {
	CalculatePay(pc paycheck.Paycheck) float64
}
