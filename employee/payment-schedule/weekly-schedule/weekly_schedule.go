package weeklySchedule

import (
	"time"

	"github.com/Rostislaved/the-payroll-system/employee/date"
)

type WeeklySchedule struct{}

func New() WeeklySchedule {
	return WeeklySchedule{}
}

func (s WeeklySchedule) IsPayday(date date.Date) bool {
	return date.Weekday() == time.Friday
}
