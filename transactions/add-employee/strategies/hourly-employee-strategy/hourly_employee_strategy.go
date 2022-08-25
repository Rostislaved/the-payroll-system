package hourlyEmployeeStrategy

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
	hourlyClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/hourly-classification"
	weeklySchedule "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-schedule/weekly-schedule"
)

type HourlyEmployeeStrategy struct {
	hourlyRate float32
}

func New(hourlyRate float32) HourlyEmployeeStrategy {
	return HourlyEmployeeStrategy{
		hourlyRate: hourlyRate,
	}
}

func (s HourlyEmployeeStrategy) MakeClassification() employee.PaymentClassification {
	return hourlyClassification.New(s.hourlyRate)
}

func (s HourlyEmployeeStrategy) MakeSchedule() employee.PaymentSchedule {
	return weeklySchedule.New()
}
