package deleteEmployee

import (
	payrollApplication "github.com/Rostislaved/the-payroll-system/payroll-application"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type DeleteEmployeeTransaction struct {
	id int
}

func DeleteEmployee(id int) payrollApplication.Transaction {
	return DeleteEmployeeTransaction{
		id: id,
	}
}

func (t DeleteEmployeeTransaction) Execute() error {
	err := payrollDatabase.DeleteEmployee(t.id)
	if err != nil {
		return err
	}

	return nil
}
