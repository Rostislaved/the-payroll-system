package addEmployeeTransaction

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-methods/holdMethod"
	payrollApplication "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-application"
	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"
)

type AddEmployeeTransaction struct {
	strategy strategy
	empid    int
	name     string
	address  string
}

type strategy interface {
	MakeClassification() employee.PaymentClassification
	MakeSchedule() employee.PaymentSchedule
}

func AddEmployee(empid int, name, address string, strategy strategy) payrollApplication.Transction {
	return AddEmployeeTransaction{
		strategy: strategy,
		empid:    empid,
		name:     name,
		address:  address,
	}
}

func (t AddEmployeeTransaction) Execute() error {
	paymentClassification := t.strategy.MakeClassification()
	paymentSchedule := t.strategy.MakeSchedule()
	paymentMethod := holdMethod.New()

	e := employee.New(t.empid, t.name, t.address)

	e.SetClassification(paymentClassification)
	e.SetSchedule(paymentSchedule)
	e.SetMethod(paymentMethod)

	err := payrollDatabase.AddEmployee(t.empid, e)
	if err != nil {
		return err
	}

	return nil
}
