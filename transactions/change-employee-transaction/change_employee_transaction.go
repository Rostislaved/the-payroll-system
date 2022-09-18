package changeEmployeeTransactions

import (
	"github.com/Rostislaved/the-payroll-system/employee"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type ChangeEmployeeTransaction struct {
	empID                  int
	changeEmployeeStrategy ChangeEmployeeStrategy
}

type ChangeEmployeeStrategy interface {
	Change(employee employee.Employee) employee.Employee
}

func ChangeEmployee(empID int, changeEmployeeStrategy ChangeEmployeeStrategy) *ChangeEmployeeTransaction {
	return &ChangeEmployeeTransaction{
		empID:                  empID,
		changeEmployeeStrategy: changeEmployeeStrategy,
	}
}

func (t ChangeEmployeeTransaction) Execute() error {
	e, err := payrollDatabase.GetEmployee(t.empID)
	if err != nil {
		return err
	}

	e = t.changeEmployeeStrategy.Change(e)

	err = payrollDatabase.DeleteEmployee(t.empID)
	if err != nil {
		return err
	}

	err = payrollDatabase.AddEmployee(t.empID, e)
	if err != nil {
		return err
	}

	return nil
}
