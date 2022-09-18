package payTransaction

import (
	"github.com/Rostislaved/the-payroll-system/employee/date"
	"github.com/Rostislaved/the-payroll-system/paycheck"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type PayTransaction struct {
	date date.Date
}

func (t PayTransaction) Execute() error {
	employeesIDs, err := payrollDatabase.GetEmployeesIDs()
	if err != nil {
		return err
	}

	for _, id := range employeesIDs {
		employee, err := payrollDatabase.GetEmployee(id)
		if err != nil {
			return err
		}

		if employee.IsPayDate(t.date) {
			pc := paycheck.New(t.date)

			employee.Payday(pc)
		}
	}

	return nil
}
