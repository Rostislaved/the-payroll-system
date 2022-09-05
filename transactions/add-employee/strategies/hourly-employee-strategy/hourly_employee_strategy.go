package hourlyEmployeeStrategy

import (
	hourlyClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification"
	weeklySchedule "github.com/Rostislaved/the-payroll-system/employee/payment-schedule/weekly-schedule"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
)

type HourlyEmployeeStrategy struct {
	hourlyRate float64
}

func New(hourlyRate float64) HourlyEmployeeStrategy {
	return HourlyEmployeeStrategy{
		hourlyRate: hourlyRate,
	}
}

func (s HourlyEmployeeStrategy) MakeClassification() paymentClassificationInterface.PaymentClassification {
	return hourlyClassification.New(s.hourlyRate)
}

func (s HourlyEmployeeStrategy) MakeSchedule() paymentScheduleInterface.PaymentSchedule {
	return weeklySchedule.New()
}
