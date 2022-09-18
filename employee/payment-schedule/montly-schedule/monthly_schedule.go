package monthlySchedule

import "github.com/Rostislaved/the-payroll-system/employee/date"

type MonthlySchedule struct{}

func New() MonthlySchedule {
	return MonthlySchedule{}
}

func (s MonthlySchedule) IsPayday(date date.Date) bool {
	m1 := date.Month()
	m2 := date.AddDays(1).Month()

	return m1 != m2
}
