package biweeklySchedule

import "github.com/Rostislaved/the-payroll-system/employee/date"

type BiweeklySchedule struct{}

func New() BiweeklySchedule {
	return BiweeklySchedule{}
}

func (s BiweeklySchedule) IsPayday(date date.Date) bool {
	return true // TODO
}

func (c BiweeklySchedule) CalculatePay(date date.Date) {
}
