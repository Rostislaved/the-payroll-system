package addEmployeeTransaction

import (
	"github.com/Rostislaved/the-payroll-system/employee"
	"github.com/Rostislaved/the-payroll-system/employee/payment-methods/holdMethod"
	"github.com/Rostislaved/the-payroll-system/interfaces/make-employee-strategy"
	payrollApplication "github.com/Rostislaved/the-payroll-system/payroll-application"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type AddEmployeeTransaction struct {
	strategy makeEmployeeStrategy.MakeEmployeeStrategy
	empID    int
	name     string
	address  string
}

func AddEmployee(empID int, name, address string, strategy makeEmployeeStrategy.MakeEmployeeStrategy) payrollApplication.Transaction {
	return AddEmployeeTransaction{
		strategy: strategy,
		empID:    empID,
		name:     name,
		address:  address,
	}
}

func (t AddEmployeeTransaction) Execute() error {
	paymentClassification := t.strategy.MakeClassification()
	paymentSchedule := t.strategy.MakeSchedule()
	paymentMethod := holdMethod.New()

	e := employee.New(t.empID, t.name, t.address)

	e.SetClassification(paymentClassification)
	e.SetSchedule(paymentSchedule)
	e.SetMethod(paymentMethod)

	err := payrollDatabase.AddEmployee(t.empID, e)
	if err != nil {
		return err
	}

	return nil
}
