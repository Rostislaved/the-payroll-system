package monthlySchedule

import "time"

type MonthlySchedule struct{}

func New() MonthlySchedule {
	return MonthlySchedule{}
}

func (s MonthlySchedule) IsPayday(date time.Time) bool {
	return true // TODO
}
