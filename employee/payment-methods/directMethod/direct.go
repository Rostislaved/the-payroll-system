package directMethod

import "fmt"

type DirectMethod struct {
	bank    string
	account string
}

func New() DirectMethod {
	return DirectMethod{}
}

func (m DirectMethod) Pay(amount int) {
	fmt.Printf("Payed %d by DirectMethod\n", amount)
}
