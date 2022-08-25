package holdMethod

import "fmt"

type HoldMethod struct{}

func New() HoldMethod {
	return HoldMethod{}
}

func (m HoldMethod) Pay(amount int) {
	fmt.Printf("Payed %d by HoldMethod\n", amount)
}
