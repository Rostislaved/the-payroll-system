package main

import (
	"os"

	payrollApplication "github.com/Rostislaved/the-payroll-system/payroll-application"
	textParserTransactionSource "github.com/Rostislaved/the-payroll-system/transaction-sources/text-parser-transaction-source"
)

func main() {
	transactionSource := textParserTransactionSource.New(os.Stdin)

	//	AddEmp 42 "Bill" "Home" H 1.23
	app := payrollApplication.New(transactionSource)

	app.Start()
}
