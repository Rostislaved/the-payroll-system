package timeCard

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
)

type TimeCard struct {
	dateTime date.Date
	hours    float64
}

func New(date date.Date, hours float64) TimeCard {
	return TimeCard{
		dateTime: date,
		hours:    hours,
	}
}

func (t *TimeCard) Date() date.Date {
	return t.dateTime
}

func (t *TimeCard) Hours() float64 {
	return t.hours
}
