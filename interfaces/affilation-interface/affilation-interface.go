package affilationInterface

import "github.com/Rostislaved/the-payroll-system/employee/date"

type Affiliation interface {
	GetFee(date date.Date) (fee float64)
}
