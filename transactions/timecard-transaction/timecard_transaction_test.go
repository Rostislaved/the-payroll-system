package timecardTransaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

	hourlyClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/hourly-classification"
	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"

	addEmployeeTransaction "my-projects/awesomeProject15_AgileSoftwareDevelopment/transactions/add-employee"
	hourlyEmployeeStrategy "my-projects/awesomeProject15_AgileSoftwareDevelopment/transactions/add-employee/strategies/hourly-employee-strategy"
)

func TestTimecardTransaction(t *testing.T) {
	payrollDatabase.Init()

	empId := 5
	d := date.New(2005, 7, 31)

	cmd := addEmployeeTransaction.AddEmployee(empId, "Bill", "Home", hourlyEmployeeStrategy.New(15.25))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	cmd = CreateTimecard(d, 8, empId)

	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	e, err := payrollDatabase.GetEmployee(empId)
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
