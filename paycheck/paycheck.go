package paycheck

import "github.com/Rostislaved/the-payroll-system/employee/date"

type Paycheck struct {
	date          date.Date
	GrossPay      float64
	DeductionsPay float64
	NetPay        float64
}

func New(date date.Date) Paycheck {
	return Paycheck{
		date: date,
	}
}

func (p *Paycheck) SetGrossPay(GrossPay float64) {
	p.GrossPay = GrossPay
}

func (p *Paycheck) SetDeductionsPay(DeductionsPay float64) {
	p.DeductionsPay = DeductionsPay
}

func (p *Paycheck) SetNetPay(NetPay float64) {
	p.NetPay = NetPay
}
