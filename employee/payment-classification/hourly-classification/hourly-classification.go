package hourlyClassification

import (
	"github.com/Rostislaved/the-payroll-system/employee/date"
	timeCard "github.com/Rostislaved/the-payroll-system/employee/payment-classification/hourly-classification/time-card"
	"github.com/Rostislaved/the-payroll-system/paycheck"

	"github.com/Rostislaved/the-payroll-system/employee"
)

type HourlyClassification struct {
	hourlyRate float64
	timeCards  []timeCard.TimeCard
}

func New(hourlyRate float64) *HourlyClassification {
	return &HourlyClassification{
		hourlyRate: hourlyRate,
	}
}

func (c *HourlyClassification) CalculatePay(pc paycheck.Paycheck) float64 {
	panic("impl")
	return 1.23
}

func (c *HourlyClassification) HourlyRate() float64 {
	return c.hourlyRate
}

func (c *HourlyClassification) GetTimeCard(date date.Date) (timeCard.TimeCard, error) {
	for _, card := range c.timeCards {
		if card.Date() == date {
			return card, nil
		}
	}

	return timeCard.TimeCard{}, employee.ErrNoTimeCard
}

func (c *HourlyClassification) AddTimeCard(card timeCard.TimeCard) {
	c.timeCards = append(c.timeCards, card)
}
