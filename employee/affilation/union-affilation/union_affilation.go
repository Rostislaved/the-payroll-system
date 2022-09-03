package unionAffilation

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
	serviceCharge "my-projects/awesomeProject15_AgileSoftwareDevelopment/transactions/add-service-charge/service-charge"
)

type UnionAffilation struct {
	charge serviceCharge.ServiceCharge
}

func New() *UnionAffilation {
	return &UnionAffilation{}
}

func (a *UnionAffilation) GetFee(date date.Date) (fee float64) {
	panic("impl")
	return 0
}

func (a *UnionAffilation) GetServiceCharge(date date.Date) serviceCharge.ServiceCharge {
	return a.charge
}

func (a *UnionAffilation) AddServiceCharge(charge serviceCharge.ServiceCharge) {
	a.charge = charge
}
