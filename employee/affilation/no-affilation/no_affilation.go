package noAffilation

import "github.com/Rostislaved/the-payroll-system/employee/date"

type NoAffilation struct{}

func New() NoAffilation {
	return NoAffilation{}
}

func (a NoAffilation) GetFee(date date.Date) (fee float64) {
	return 0
}
