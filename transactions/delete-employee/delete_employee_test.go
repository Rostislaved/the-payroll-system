package deleteEmployee

import (
	"errors"
	"testing"

	"github.com/Rostislaved/the-payroll-system/employee"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"

	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
	commissionedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/commissioned-employee-strategy"
)

func TestDeleteEmployee(t *testing.T) {
	payrollDatabase.Init()

	empID := 4

	cs := commissionedEmployeeStrategy.New(2500, 3.2)

	at := addEmployeeTransaction.AddEmployee(empID, "Bill", "Home", cs)

	err := at.Execute()
	if err != nil {
		t.Fatal(err)
	}

	_, err = payrollDatabase.GetEmployee(empID)
	if err != nil {
		t.Fatal(err)
	}

	dt := DeleteEmployee(empID)

	err = dt.Execute()
	if err != nil {
		t.Fatal()
	}

	_, err = payrollDatabase.GetEmployee(empID)
	if err != nil {
		switch {
		case errors.Is(err, employee.ErrEmployeeNotFound):

		default:
			t.Fatal(err)
		}
	} else {
		t.Fatal()
	}
}

//{
//  int empID = 4;
//  AddCommissionedEmployee t =
//    new AddCommissionedEmployee(
//      empID, “Bill”, “Home”, 2500, 3.2);
//  t.Execute();
//  Employee e = PayrollDatabase.GetEmployee(empID);
//  Assert.IsNotNull(e);
//  DeleteEmployeeTransaction dt =
//    new DeleteEmployeeTransaction(empID);
//  dt.Execute();
//  e = PayrollDatabase.GetEmployee(empID);
//  Assert.IsNull(e);
//}
