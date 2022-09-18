package salesReceiptTransaction

import (
	"errors"

	salesReceipt "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification/sales-receipt"

	"github.com/Rostislaved/the-payroll-system/employee/date"
	commissionedClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/commissioned-classification"
	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"
)

type SalesReceiptTransaction struct {
	date   date.Date
	amount float64
	empID  int
}

func CreateSalesReceipt(date date.Date, amount float64, empID int) *SalesReceiptTransaction {
	return &SalesReceiptTransaction{
		date:   date,
		amount: amount,
		empID:  empID,
	}
}

func (t SalesReceiptTransaction) Execute() error {
	e, err := payrollDatabase.GetEmployee(t.empID)
	if err != nil {
		return err
	}

	classification := e.Classification()

	cc, ok := classification.(*commissionedClassification.CommissionedClassification)
	if !ok {
		return errors.New("tried to add sales receipt to non-commissioned employee")
	}

	sr := salesReceipt.New(t.amount, t.date)

	cc.AddSalesReceipt(sr)

	return nil
}
