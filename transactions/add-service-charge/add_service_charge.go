package addServiceCharge

import (
	"errors"

	unionAffiliation "github.com/Rostislaved/the-payroll-system/employee/affiliation/union-affiliation"
	"github.com/Rostislaved/the-payroll-system/employee/date"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
	serviceCharge "github.com/Rostislaved/the-payroll-system/transactions/add-service-charge/service-charge"
)

type ServiceChargeTransaction struct {
	memberID int
	date     date.Date
	charge   float64
}

func CreateServiceCharge(memberID int, date date.Date, charge float64) *ServiceChargeTransaction {
	return &ServiceChargeTransaction{
		memberID: memberID,
		date:     date,
		charge:   charge,
	}
}

func (t ServiceChargeTransaction) Execute() error {
	employee, err := payrollDatabase.GetUnionMember(t.memberID)
	if err != nil {
		return err
	}

	ua, ok := employee.Affiliation().(*unionAffiliation.UnionAffiliation)
	if !ok {
		err = errors.New("tries to add service charge to union member without a union affiliation")

		return err
	}

	ua.AddServiceCharge(serviceCharge.New(t.date, t.charge))

	return nil
}
