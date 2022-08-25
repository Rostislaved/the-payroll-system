package payrollApplication

import payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"

type PayrollApplication struct {
	ts TransactionSource
}

type TransactionSource interface {
	GetTransaction() Transction
}
type Transction interface {
	Execute()
}

func New(ts TransactionSource) PayrollApplication {
	payrollDatabase.Init()

	return PayrollApplication{ts: ts}
}

func (pa *PayrollApplication) Start() {
	for {
		t := pa.ts.GetTransaction()

		t.Execute()
	}
}
