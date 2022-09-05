package paymentScheduleInterface

import "github.com/Rostislaved/the-payroll-system/employee/date"

type PaymentSchedule interface {
	IsPayday(date date.Date) bool
}
