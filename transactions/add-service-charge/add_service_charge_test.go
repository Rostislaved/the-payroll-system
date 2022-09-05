package addServiceCharge

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Rostislaved/the-payroll-system/employee/date"

	unionAffilation "github.com/Rostislaved/the-payroll-system/employee/affilation/union-affilation"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"

	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
	hourlyEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/hourly-employee-strategy"
)

func TestServiceChargeTransaction(t *testing.T) {
	payrollDatabase.Init()

	empID := 2

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bill", "Home", hourlyEmployeeStrategy.New(15.25))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	employee, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	ua := unionAffilation.New()

	employee.SetAffiliation(ua)

	memberID := 86

	err = payrollDatabase.AddUnionMember(memberID, employee)
	if err != nil {
		t.Fatal(err)
	}

	date := date.New(2005, 8, 8)
	charge := 12.95

	cmd = CreateServiceCharge(memberID, date, charge)

	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	serviceCharge := ua.GetServiceCharge(date)

	assert.Equal(t, charge, serviceCharge.Amount(), 0.001)
}
