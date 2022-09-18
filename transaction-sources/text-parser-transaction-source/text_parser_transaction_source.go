package textParserTransactionSource

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	commissionedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/commissioned-employee-strategy"

	salariedEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/salaried-employee-strategy"

	makeEmployeeStrategy "github.com/Rostislaved/the-payroll-system/interfaces/make-employee-strategy"

	addEmployeeTransaction "github.com/Rostislaved/the-payroll-system/transactions/add-employee"
	hourlyEmployeeStrategy "github.com/Rostislaved/the-payroll-system/transactions/add-employee/strategies/hourly-employee-strategy"

	payrollApplication "github.com/Rostislaved/the-payroll-system/payroll-application"
)

type TextParserTransactionSource struct {
	scanner *bufio.Scanner
}

func New(reader io.Reader) *TextParserTransactionSource {
	scanner := bufio.NewScanner(os.Stdin)

	return &TextParserTransactionSource{
		scanner: scanner,
	}
}

func (ts *TextParserTransactionSource) GetTransaction() (payrollApplication.Transaction, error) {
	var line string
	if ts.scanner.Scan() {
		line = ts.scanner.Text()
	}

	if ts.scanner.Err() != nil {
		return nil, ts.scanner.Err()
	}

	t, err := parseLine(line)
	if err != nil {
		return nil, err
	}

	return t, nil
}

var ErrIncorrectInput = errors.New("input line is incorrect")

func parseLine(line string) (transction payrollApplication.Transaction, err error) {
	tokens := strings.Split(line, " ")

	if len(tokens) <= 1 {
		return nil, ErrIncorrectInput
	}

	transactionName := tokens[0]

	switch transactionName {
	case "AddEmp":
		transction, err = parseAddEmployee(tokens)
		if err != nil {
			return nil, err
		}

	case "DelEmp":
		panic("Not implemented")
	case "TimeCard":
		panic("Not implemented")
	case "SalesReceipt":
		panic("Not implemented")
	case "ServiceCharge":
		panic("Not implemented")
	case "ChgEmp":
		panic("Not implemented")
	case "Payday":
		panic("Not implemented")
	default:
		return nil, ErrIncorrectInput

	}

	return transction, nil
}

func parseAddEmployee(tokens []string) (transaction payrollApplication.Transaction, err error) {
	if len(tokens) < 5 {
		return nil, ErrIncorrectInput
	}

	empID, err := strconv.Atoi(tokens[1])
	if err != nil {
		return nil, ErrIncorrectInput
	}

	name := tokens[2]
	address := tokens[3]

	classification := tokens[4]

	var strategy makeEmployeeStrategy.MakeEmployeeStrategy

	switch classification {
	case "H":
		if len(tokens) != 6 {
			return nil, ErrIncorrectInput
		}

		hourlyRate, err := strconv.ParseFloat(tokens[5], 64)
		if err != nil {
			return nil, ErrIncorrectInput
		}

		strategy = hourlyEmployeeStrategy.New(hourlyRate)
	case "S":
		if len(tokens) != 6 {
			return nil, ErrIncorrectInput
		}

		monthlyRate, err := strconv.ParseFloat(tokens[5], 64)
		if err != nil {
			return nil, ErrIncorrectInput
		}

		strategy = salariedEmployeeStrategy.New(monthlyRate)
	case "C":
		if len(tokens) != 7 {
			return nil, ErrIncorrectInput
		}

		monthlyRate, err := strconv.ParseFloat(tokens[5], 64)
		if err != nil {
			return nil, ErrIncorrectInput
		}

		commissionRate, err := strconv.ParseFloat(tokens[6], 64)
		if err != nil {
			return nil, ErrIncorrectInput
		}

		strategy = commissionedEmployeeStrategy.New(monthlyRate, commissionRate)
	default:
		return nil, ErrIncorrectInput

	}

	transaction = addEmployeeTransaction.AddEmployee(empID, name, address, strategy)

	return transaction, nil
}
