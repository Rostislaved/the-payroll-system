package commissionedEmployeeStrategy

import (
	commissionedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification"
	biweeklySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/biweekly-schedule"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
)

type CommissionedEmployeeStrategy struct {
	salary         float64
	commissionRate float64
}

func New(salary, commissionRate float64) CommissionedEmployeeStrategy {
	return CommissionedEmployeeStrategy{
		salary:         salary,
		commissionRate: commissionRate,
	}
}

func (s CommissionedEmployeeStrategy) MakeClassification() paymentClassificationInterface.PaymentClassification {
	return commissionedClassification.New(s.salary, s.commissionRate)
}

func (s CommissionedEmployeeStrategy) MakeSchedule() paymentScheduleInterface.PaymentSchedule {
	return biweeklySchedule.New()
}
