package commissionedEmployeeStrategy

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
	commissionedClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/commissioned-classification"
	biweeklySchedule "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-schedule/biweekly-schedule"
)

type CommissionedEmployeeStrategy struct {
	salary         float32
	commissionRate float32
}

func New(salary, commissionRate float32) CommissionedEmployeeStrategy {
	return CommissionedEmployeeStrategy{
		salary:         salary,
		commissionRate: commissionRate,
	}
}

func (s CommissionedEmployeeStrategy) MakeClassification() employee.PaymentClassification {
	return commissionedClassification.New(s.salary, s.commissionRate)
}

func (s CommissionedEmployeeStrategy) MakeSchedule() employee.PaymentSchedule {
	return biweeklySchedule.New()
}
