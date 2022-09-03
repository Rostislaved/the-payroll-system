package salesReceiptTransaction

import (
	"errors"

	salesReceipt "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/commissioned-classification/sales-receipt"

	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
	commissionedClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/commissioned-classification"
	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"
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
