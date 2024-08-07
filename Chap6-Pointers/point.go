package pointers

import "fmt"

/*In go if a symbol (variables,types,functions,etc) starts with a
lowercase then it is private outisde of that package*/

type Bitcoin int

type Wallet struct {
	money Bitcoin
}

//Native interface of the fmt package
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*when you call a function or method the arguments are copied
since you need to reference the exact same wallet to guarantee
the money is altered you need to pass a pointer to the struct you wish to alter*/
func (w *Wallet) Deposit(amount Bitcoin) {
	w.money += amount
}

/*using a pointer here is unnecessery since returning
a copy would be just fine but we do it for consistency*/
func (w *Wallet) Balance() Bitcoin {
	return w.money
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.money -= amount
}
