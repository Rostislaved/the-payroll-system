package timecardTransaction

import (
	"errors"

	"github.com/Rostislaved/the-payroll-system/employee/date"
	timeCard "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification/time-card"

	hourlyClassification "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification"

	payrollDatabase "github.com/Rostislaved/the-payroll-system/payroll-database"

	payrollApplication "github.com/Rostislaved/the-payroll-system/payroll-application"
)

type TimecardTransaction struct {
	date  date.Date
	hours float64
	empID int
}

func CreateTimecard(date date.Date, hours float64, empID int) payrollApplication.Transction {
	return TimecardTransaction{
		date:  date,
		hours: hours,
		empID: empID,
	}
}

func (t TimecardTransaction) Execute() error {
	e, err := payrollDatabase.GetEmployee(t.empID)
	if err != nil {
		return err
	}

	classification := e.Classification()

	hc, ok := classification.(*hourlyClassification.HourlyClassification)
	if !ok {
		return errors.New("tried to add timecard to non-hourly employee")
	}

	hc.AddTimeCard(timeCard.New(t.date, t.hours))

	return nil
}
