package salariedEmployeeStrategy

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
	salariedClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/salaried-classification"
	monthlySchedule "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-schedule/montly-schedule"
)

type SalariedEmployeeStrategy struct {
	salary float32
}

func New(salary float32) SalariedEmployeeStrategy {
	return SalariedEmployeeStrategy{
		salary: salary,
	}
}

func (s SalariedEmployeeStrategy) MakeClassification() employee.PaymentClassification {
	return salariedClassification.New(s.salary)
}

func (s SalariedEmployeeStrategy) MakeSchedule() employee.PaymentSchedule {
	return monthlySchedule.New()
}
