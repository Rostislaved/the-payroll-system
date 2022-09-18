package biweeklySchedule

import (
	"time"

	"github.com/Rostislaved/the-payroll-system/employee/date"
)

type BiweeklySchedule struct {
	secondFriday bool
}

func New() BiweeklySchedule {
	return BiweeklySchedule{}
}

func (s BiweeklySchedule) IsPayday(date date.Date) bool {
	if date.Weekday() == time.Friday {
		if s.secondFriday {
			s.secondFriday = false
			return true
		} else {
			s.secondFriday = true
			return false
		}
	}

	return false
}
