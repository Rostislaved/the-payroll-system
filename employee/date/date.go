package date

import "time"

type Date time.Time

func New(year, month, day int) Date {
	return Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}
