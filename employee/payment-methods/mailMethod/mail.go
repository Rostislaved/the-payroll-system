package mailMethod

import "fmt"

type MailMethod struct {
	address string
}

func New() MailMethod {
	return MailMethod{}
}

func (m MailMethod) Pay(amount int) {
	fmt.Printf("Payed %d by MailMethod\n", amount)
}
