package biweeklySchedule

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type BiweeklySchedule struct{}

func New() BiweeklySchedule {
	return BiweeklySchedule{}
}

func (s BiweeklySchedule) IsPayday(date date.Date) bool {
	return true // TODO
}

func (c BiweeklySchedule) CalculatePay(date date.Date) {
}
