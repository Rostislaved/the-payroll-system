package holdMethod

import (
	"github.com/Rostislaved/the-payroll-system/paycheck"
)

type HoldMethod struct{}

func New() HoldMethod {
	return HoldMethod{}
}

func (m HoldMethod) Pay(pc paycheck.Paycheck) {
	panic("impl")
}
