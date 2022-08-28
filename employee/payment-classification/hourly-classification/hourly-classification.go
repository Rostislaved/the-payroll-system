package hourlyClassification

import (
	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/date"

	timeCard "my-projects/awesomeProject15_AgileSoftwareDevelopment/employee/payment-classification/hourly-classification/time-card"

	"my-projects/awesomeProject15_AgileSoftwareDevelopment/employee"
)

var _ employee.PaymentClassification = (*HourlyClassification)(nil)

type HourlyClassification struct {
	hourlyRate float64
	timeCards  []timeCard.TimeCard
}

func New(hourlyRate float64) *HourlyClassification {
	return &HourlyClassification{
		hourlyRate: hourlyRate,
	}
}

func (c *HourlyClassification) CalculatePay(date date.Date) float64 {
	return 1.23 // impl
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
