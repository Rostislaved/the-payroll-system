package noAffilation

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type NoAffilation struct{}

func New() NoAffilation {
	return NoAffilation{}
}

func (a NoAffilation) GetFee(date date.Date) (fee float64) {
	return 0
}
