package unionAffiliation

import (
	"github.com/Rostislaved/the-payroll-system/employee/date"
	serviceCharge "github.com/Rostislaved/the-payroll-system/transactions/add-service-charge/service-charge"
)

type UnionAffiliation struct {
	charge serviceCharge.ServiceCharge
}

func New() *UnionAffiliation {
	return &UnionAffiliation{}
}

func (a *UnionAffiliation) GetFee(date date.Date) (fee float64) {
	panic("impl")
	return 0
}

func (a *UnionAffiliation) GetServiceCharge(date date.Date) serviceCharge.ServiceCharge {
	return a.charge
}

func (a *UnionAffiliation) AddServiceCharge(charge serviceCharge.ServiceCharge) {
	a.charge = charge
}
