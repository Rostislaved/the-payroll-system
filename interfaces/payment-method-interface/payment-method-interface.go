package paymentMethodInterface

import "github.com/Rostislaved/the-payroll-system/paycheck"

type PaymentMethod interface {
	Pay(pc paycheck.Paycheck)
}
