package serviceCharge

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type ServiceCharge struct {
	date   date.Date
	amount float64
}

func New(date date.Date, amount float64) ServiceCharge {
	return ServiceCharge{
		date:   date,
		amount: amount,
	}
}

func (c ServiceCharge) Amount() float64 {
	return c.amount
}
