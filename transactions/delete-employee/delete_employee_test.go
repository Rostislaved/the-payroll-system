package deleteEmployee

import (
	"errors"
	"testing"

	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"

	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"

	addEmployeeTransaction "my-projects/awesomeProject15_AgileSoftwareDevelopment/transactions/add-employee"
	commissionedEmployeeStrategy "my-projects/awesomeProject15_AgileSoftwareDevelopment/transactions/add-employee/strategies/commissioned-employee-strategy"
)

func TestDeleteEmployee(t *testing.T) {
	payrollDatabase.Init()

	empId := 4

	cs := commissionedEmployeeStrategy.New(2500, 3.2)

	at := addEmployeeTransaction.AddEmployee(empId, "Bill", "Home", cs)

	err := at.Execute()
	if err != nil {
		t.Fatal(err)
	}

	_, err = payrollDatabase.GetEmployee(empId)
	if err != nil {
		t.Fatal(err)
	}

	dt := DeleteEmployee(empId)

	err = dt.Execute()
	if err != nil {
		t.Fatal()
	}

	_, err = payrollDatabase.GetEmployee(empId)
	if err != nil {
		switch {
		case errors.Is(err, employee.ErrNoEmployeeFound):

		default:
			t.Fatal(err)
		}
	} else {
		t.Fatal()
	}
}

//{
//  int empId = 4;
//  AddCommissionedEmployee t =
//    new AddCommissionedEmployee(
//      empId, “Bill”, “Home”, 2500, 3.2);
//  t.Execute();
//  Employee e = PayrollDatabase.GetEmployee(empId);
//  Assert.IsNotNull(e);
//  DeleteEmployeeTransaction dt =
//    new DeleteEmployeeTransaction(empId);
//  dt.Execute();
//  e = PayrollDatabase.GetEmployee(empId);
//  Assert.IsNull(e);
//}
