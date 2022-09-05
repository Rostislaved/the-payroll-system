package salesReceipt

import "github.com/Rostislaved/the-payroll-system/employee/date"

type SalesReceipt struct {
	amount float64
	date   date.Date
}

func New(amount float64, date date.Date) SalesReceipt {
	return SalesReceipt{
		amount: amount,
		date:   date,
	}
}

func (s SalesReceipt) Amount() float64 {
	return s.amount
}

func (s SalesReceipt) Date() date.Date {
	return s.date
}
