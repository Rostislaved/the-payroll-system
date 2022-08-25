package hourlyClassification

import "time"

type TimeCard struct {
	// implement
}

type HourlyClassification struct {
	hourlyRate float32
	timeCards  []TimeCard
}

func New(hourlyRate float32) HourlyClassification {
	return HourlyClassification{
		hourlyRate: hourlyRate,
	}
}

func (c HourlyClassification) CalculatePay(date time.Time) float32 {
	return 1.23 // impl
}

func (c HourlyClassification) HourlyRate() float32 {
	return c.hourlyRate
}
