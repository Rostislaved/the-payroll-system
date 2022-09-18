package changeEmployeeTransactions

import (
	"testing"

	"github.com/Rostislaved/the-payroll-system/employee/payment-methods/directMethod"

	changeMethodStrategy "github.com/Rostislaved/the-payroll-system/transactions/change-employee-transaction/strategies/change-method-strategy"

	hourlyEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/hourly-employee-strategy"
	changeClassificationStrategy "github.com/Rostislaved/the-payroll-system/transactions/change-employee-transaction/strategies/change-classification-strategy"

	changeAddressStrategy "github.com/Rostislaved/the-payroll-system/transactions/change-employee-transaction/strategies/change-address-strategy"

	"github.com/stretchr/testify/assert"

	changeNameStrategy "github.com/Rostislaved/the-payroll-system/transactions/change-employee-transaction/strategies/change-name-strategy"

	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
	salariedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/salaried-employee-strategy"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

func TestChangeName(t *testing.T) {
	payrollDatabase.Init()

	strategy := salariedEmployeeStrategy.New(1.23)

	empID := 42

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bob", "Home", strategy)

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	changeStrategy := changeNameStrategy.New("Mike")

	cmdChangeEmployee := ChangeEmployee(empID, changeStrategy)

	err = cmdChangeEmployee.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Mike", e.Name())
}

func TestChangeAddress(t *testing.T) {
	payrollDatabase.Init()

	strategy := salariedEmployeeStrategy.New(1.23)

	empID := 42

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bob", "Home", strategy)

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	changeStrategy := changeAddressStrategy.New("Work")

	cmdChangeEmployee := ChangeEmployee(empID, changeStrategy)

	err = cmdChangeEmployee.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Work", e.Address())
}

func TestChangeClassification(t *testing.T) {
	payrollDatabase.Init()

	salariedStrategy := salariedEmployeeStrategy.New(1.23)

	empID := 42

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bob", "Home", salariedStrategy)

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	hourlyStrategy := hourlyEmployeeStrategy.New(1.23)

	changeStrategy := changeClassificationStrategy.New(hourlyStrategy)

	cmdChangeEmployee := ChangeEmployee(empID, changeStrategy)

	err = cmdChangeEmployee.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, hourlyStrategy.MakeClassification(), e.Classification())
}

func TestChangeMethod(t *testing.T) {
	payrollDatabase.Init()

	salariedStrategy := salariedEmployeeStrategy.New(1.23)

	empID := 42

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bob", "Home", salariedStrategy)

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	method := directMethod.New()

	changeStrategy := changeMethodStrategy.New(method)

	cmdChangeEmployee := ChangeEmployee(empID, changeStrategy)

	err = cmdChangeEmployee.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, method, e.Method())
}
