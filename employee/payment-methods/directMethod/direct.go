package directMethod

import (
	"github.com/Rostislaved/the-payroll-system/paycheck"
)

type DirectMethod struct {
	bank    string
	account string
}

func New() DirectMethod {
	return DirectMethod{}
}

func (m DirectMethod) Pay(pc paycheck.Paycheck) {
	panic("impl")
}
