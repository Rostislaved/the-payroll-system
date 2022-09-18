package affiliationInterface

import (
	"github.com/Rostislaved/the-payroll-system/paycheck"
)

type Affiliation interface {
	GetFee(pc paycheck.Paycheck) (fee float64) // todo CalculateDeductions
}
