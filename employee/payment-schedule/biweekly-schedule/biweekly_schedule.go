package biweeklySchedule

import "time"

type BiweeklySchedule struct{}

func New() BiweeklySchedule {
	return BiweeklySchedule{}
}

func (s BiweeklySchedule) IsPayday(date time.Time) bool {
	return true // TODO
}

func (c BiweeklySchedule) CalculatePay(date time.Time) {
}
