package payrollApplication

import (
	"log"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type PayrollApplication struct {
	ts TransactionSource
}

type TransactionSource interface {
	GetTransaction() (Transaction, error)
}
type Transaction interface {
	Execute() error
}

func New(ts TransactionSource) PayrollApplication {
	payrollDatabase.Init()

	return PayrollApplication{
		ts: ts,
	}
}

func (pa *PayrollApplication) Start() {
	for {
		t, err := pa.ts.GetTransaction()
		if err != nil {
			log.Println(err)

			continue
		}

		err = t.Execute()
		if err != nil {
			log.Println(err)

			continue
		}
	}
}
