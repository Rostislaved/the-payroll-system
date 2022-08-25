package weeklySchedule

import "time"

type WeeklySchedule struct{}

func New() WeeklySchedule {
	return WeeklySchedule{}
}

func (s WeeklySchedule) IsPayday(date time.Time) bool {
	return true // TODO
}
