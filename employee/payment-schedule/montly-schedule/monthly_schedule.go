package monthlySchedule

import "github.com/Rostislaved/the-payroll-system/employee/date"

type MonthlySchedule struct{}

func New() MonthlySchedule {
	return MonthlySchedule{}
}

func (s MonthlySchedule) IsPayday(date date.Date) bool {
	return true // TODO
}
