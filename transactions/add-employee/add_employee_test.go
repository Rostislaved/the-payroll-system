package addEmployeeTransaction

import (
	"testing"

	biweeklySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/biweekly-schedule"

	commissionedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification"
	commissionedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/commissioned-employee-strategy"

	hourlyEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/hourly-employee-strategy"

	hourlyClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification"
	weeklySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/weekly-schedule"

	"github.com/Rostislaved/the-payroll-system/employee/payment-methods/holdMethod"

	monthlySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/montly-schedule"

	salariedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/salaried-classification"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
	salariedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/salaried-employee-strategy"
	"github.com/stretchr/testify/assert"
)

var (
	empID   = 1
	name    = "Bob"
	address = "Home"
)

func TestAddSalariedEmployee(t *testing.T) {
	payrollDatabase.Init()

	salary := 1000.0

	cmd := AddEmployee(empID, name, address, salariedEmployeeStrategy.New(salary))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	employee, err := payrollDatabase.GetEmployee(1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, name, employee.Name())

	assert.IsType(t, salariedClassification.SalariedClassification{}, employee.Classification())

	sc, _ := employee.Classification().(salariedClassification.SalariedClassification)
	assert.InDelta(t, salary, sc.Salary(), 0.001)

	assert.IsType(t, monthlySchedule.MonthlySchedule{}, employee.Schedule())
	assert.IsType(t, holdMethod.HoldMethod{}, employee.Method())
}

func TestAddHourlyEmployee(t *testing.T) {
	payrollDatabase.Init()

	hourlyRate := 10.42

	cmd := AddEmployee(empID, name, address, hourlyEmployeeStrategy.New(hourlyRate))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	employee, err := payrollDatabase.GetEmployee(1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, name, employee.Name())

	assert.IsType(t, &hourlyClassification.HourlyClassification{}, employee.Classification())

	hc, _ := employee.Classification().(*hourlyClassification.HourlyClassification)
	assert.InDelta(t, hourlyRate, hc.HourlyRate(), 0.001)

	assert.IsType(t, weeklySchedule.WeeklySchedule{}, employee.Schedule())
	assert.IsType(t, holdMethod.HoldMethod{}, employee.Method())
}

func TestAddCommissionedEmployee(t *testing.T) {
	payrollDatabase.Init()

	salary := 1000.42
	commissionRate := 10.42

	cmd := AddEmployee(empID, name, address, commissionedEmployeeStrategy.New(salary, commissionRate))

	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	employee, err := payrollDatabase.GetEmployee(1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, name, employee.Name())

	assert.IsType(t, &commissionedClassification.CommissionedClassification{}, employee.Classification())

	cc, _ := employee.Classification().(*commissionedClassification.CommissionedClassification)
	assert.InDelta(t, salary, cc.Salary(), 0.001)
	assert.InDelta(t, commissionRate, cc.CommissionRate(), 0.001)

	assert.IsType(t, biweeklySchedule.BiweeklySchedule{}, employee.Schedule())
	assert.IsType(t, holdMethod.HoldMethod{}, employee.Method())
}
