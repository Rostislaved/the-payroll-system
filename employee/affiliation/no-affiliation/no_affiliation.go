package noAffiliation

import "github.com/Rostislaved/the-payroll-system/employee/date"

type NoAffiliation struct{}

func New() NoAffiliation {
	return NoAffiliation{}
}

func (a NoAffiliation) GetFee(date date.Date) (fee float64) {
	return 0
}
