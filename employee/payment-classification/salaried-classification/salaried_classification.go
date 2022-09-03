package salariedClassification

import "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

type SalariedClassification struct {
	salary float64
}

func New(salary float64) SalariedClassification {
	return SalariedClassification{
		salary: salary,
	}
}

func (c SalariedClassification) Salary() float64 {
	return c.salary
}

func (c SalariedClassification) CalculatePay(date date.Date) float64 {
	return c.salary
}
