package salariedClassification

import "time"

type SalariedClassification struct {
	salary float32
}

func New(salary float32) SalariedClassification {
	return SalariedClassification{
		salary: salary,
	}
}

func (c SalariedClassification) Salary() float32 {
	return c.salary
}

func (c SalariedClassification) CalculatePay(date time.Time) float32 {
	return 1.23 // impl
}
