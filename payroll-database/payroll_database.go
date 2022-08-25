package payrollDatabase

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"

type PayrollDatabase struct {
	employees map[int]employee.Employee
}

var PayrollDatabaseInstance = PayrollDatabase{}

func Init() {
	PayrollDatabaseInstance.employees = make(map[int]employee.Employee)
}

func AddEmployee(id int, employee employee.Employee) {
	PayrollDatabaseInstance.employees[id] = employee
}

func GetEmployee(id int) (employee employee.Employee) {
	return PayrollDatabaseInstance.employees[id]
}
