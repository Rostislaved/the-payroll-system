package payrollDatabase

import (
	"errors"

	"github.com/Rostislaved/the-payroll-system/employee"
)

type PayrollDatabase struct {
	employees map[int]employee.Employee
	members   map[int]employee.Employee
}

var PayrollDatabaseInstance = PayrollDatabase{}

func Init() {
	PayrollDatabaseInstance.employees = make(map[int]employee.Employee)
	PayrollDatabaseInstance.members = make(map[int]employee.Employee)
}

func AddEmployee(id int, employee employee.Employee) error {
	PayrollDatabaseInstance.employees[id] = employee

	return nil
}

func GetEmployee(id int) (e employee.Employee, err error) {
	e, ok := PayrollDatabaseInstance.employees[id]
	if !ok {
		err = employee.ErrEmployeeNotFound

		return employee.Employee{}, err
	}

	return e, nil
}

func GetEmployeesIDs() (employeesIDs []int, err error) {
	for id := range PayrollDatabaseInstance.employees {
		employeesIDs = append(employeesIDs, id)
	}

	return employeesIDs, nil
}

func DeleteEmployee(id int) error {
	delete(PayrollDatabaseInstance.employees, id)

	return nil
}

func AddUnionMember(memberID int, e employee.Employee) error {
	PayrollDatabaseInstance.members[memberID] = e

	return nil
}

func GetUnionMember(memberID int) (employee.Employee, error) {
	e, ok := PayrollDatabaseInstance.members[memberID]
	if !ok {
		err := errors.New("union member not found")

		return employee.Employee{}, err
	}

	return e, nil
}
