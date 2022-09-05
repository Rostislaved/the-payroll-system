package timecardTransaction

import (
	"testing"

	"github.com/Rostislaved/the-payroll-system/employee/date"
	"github.com/stretchr/testify/assert"

	hourlyClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"

	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
	hourlyEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/hourly-employee-strategy"
)

func TestTimecardTransaction(t *testing.T) {
	payrollDatabase.Init()

	empID := 5
	d := date.New(2005, 7, 31)

	cmd := addEmployeeTransaction.AddEmployee(empID, "Bill", "Home", hourlyEmployeeStrategy.New(15.25))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	cmd = CreateTimecard(d, 8, empID)

	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	classification := e.Classification()

	hc, ok := classification.(*hourlyClassification.HourlyClassification)
	if !ok {
		t.Fatal()
	}

	timecard, err := hc.GetTimeCard(d)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 8.0, timecard.Hours())
}
