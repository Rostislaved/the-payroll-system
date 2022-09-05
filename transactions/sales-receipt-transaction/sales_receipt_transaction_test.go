package salesReceiptTransaction

import (
	"testing"

	salesReceipt "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification/sales-receipt"
	"github.com/stretchr/testify/assert"

	"github.com/Rostislaved/the-payroll-system/employee/date"

	commissionedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification"
	commissionedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/commissioned-employee-strategy"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
)

//func a()  {
//	payrollDatabase.Init()
//
//	empID := 5
//	d := date.New(2005, 7, 31)
//
//	cmd := addEmployeeTransaction.AddEmployee(empID, "Bill", "Home", hourlyEmployeeStrategy.New(15.25))
//
//	err := cmd.Execute()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	cmd = CreateTimecard(d, 8, empID)
//
//	err = cmd.Execute()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	e, err := payrollDatabase.GetEmployee(empID)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	classification := e.Classification()
//
//	hc, ok := classification.(*hourlyClassification.HourlyClassification)
//	if !ok {
//		t.Fatal()
//	}
//
//	timecard, err := hc.GetTimeCard(d)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	assert.Equal(t, 8.0, timecard.Hours())
//}

func TestSalesReceiptTransaction(t *testing.T) {
	payrollDatabase.Init()

	empID := 1
	name := "Bob"
	address := "Home"

	salary := 1000.42
	commissionRate := 10.42

	cmd := addEmployeeTransaction.AddEmployee(empID, name, address, commissionedEmployeeStrategy.New(salary, commissionRate))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	date := date.New(2005, 7, 31)
	amount := 42.42

	cmd = CreateSalesReceipt(date, amount, empID)

	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	employee, err := payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	cc, ok := employee.Classification().(*commissionedClassification.CommissionedClassification)
	if !ok {
		t.Fatal()
	}

	salesReceipts := cc.SalesReceipts()

	assert.Contains(t, salesReceipts, salesReceipt.New(amount, date))
}
