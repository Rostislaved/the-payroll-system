package timecardTransaction

import (
	"errors"

	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"
	timeCard "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/hourly-classification/time-card"

	hourlyClassification "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/hourly-classification"

	payrollDatabase "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-database"

	payrollApplication "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-application"
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
