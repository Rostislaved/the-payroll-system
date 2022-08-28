package monthlySchedule

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type MonthlySchedule struct{}

func New() MonthlySchedule {
	return MonthlySchedule{}
}

func (s MonthlySchedule) IsPayday(date date.Date) bool {
	return true // TODO
}
