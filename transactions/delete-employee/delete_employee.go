package deleteEmployee

import (
	payrollApplication "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-application"
	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"
)

type DeleteEmployeeTransaction struct {
	id int
}

func DeleteEmployee(id int) payrollApplication.Transction {
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
