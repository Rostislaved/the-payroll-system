package payrollDatabase

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
)

type PayrollDatabase struct {
	employees map[int]employee.Employee
}

var PayrollDatabaseInstance = PayrollDatabase{}

func Init() {
	PayrollDatabaseInstance.employees = make(map[int]employee.Employee)
}

func AddEmployee(id int, employee employee.Employee) error {
	PayrollDatabaseInstance.employees[id] = employee

	return nil
}

func GetEmployee(id int) (e employee.Employee, err error) {
	e, ok := PayrollDatabaseInstance.employees[id]
	if !ok {
		err = employee.ErrNoEmployeeFound

		return employee.Employee{}, err
	}

	return e, nil
}

func DeleteEmployee(id int) error {
	delete(PayrollDatabaseInstance.employees, id)

	return nil
}
