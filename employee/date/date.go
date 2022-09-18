package date

import "time"

type Date struct {
	time.Time
}

func New(year, month, day int) Date {
	return Date{
		Time: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
	}
}

func (d *Date) Month() int {
	return d.Month()
}

func (d *Date) AddDays(numberOfDays int) *Date {
	d.AddDate(0, 0, numberOfDays)

	return d
}
