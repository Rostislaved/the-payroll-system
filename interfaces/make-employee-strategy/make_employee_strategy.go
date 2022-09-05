package makeEmployeeStrategy

import (
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-classification-interface"
	"github.com/Rostislaved/the-payroll-system/interfaces/payment-schedule-interface"
)

type MakeEmployeeStrategy interface {
	MakeClassification() paymentClassificationInterface.PaymentClassification
	MakeSchedule() paymentScheduleInterface.PaymentSchedule
}
