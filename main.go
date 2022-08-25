package main

import payrollApplication "my-projects/awesomeProject15_AgileSoftwareDevelopment/payroll-application"

func main() {
	var ts payrollApplication.TransactionSource

	app := payrollApplication.New(ts)

	app.Start()
}
