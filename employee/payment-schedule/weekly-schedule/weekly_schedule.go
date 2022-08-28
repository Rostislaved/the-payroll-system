package weeklySchedule

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type WeeklySchedule struct{}

func New() WeeklySchedule {
	return WeeklySchedule{}
}

func (s WeeklySchedule) IsPayday(date date.Date) bool {
	return true // TODO
}
