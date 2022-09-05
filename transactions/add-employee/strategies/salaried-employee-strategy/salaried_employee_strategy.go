package salariedEmployeeStrategy

import (
	salariedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/salaried-classification"
	monthlySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/montly-schedule"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
)

type SalariedEmployeeStrategy struct {
	salary float64
}

func New(salary float64) SalariedEmployeeStrategy {
	return SalariedEmployeeStrategy{
		salary: salary,
	}
}

func (s SalariedEmployeeStrategy) MakeClassification() paymentClassificationInterface.PaymentClassification {
	return salariedClassification.New(s.salary)
}

func (s SalariedEmployeeStrategy) MakeSchedule() paymentScheduleInterface.PaymentSchedule {
	return monthlySchedule.New()
}
