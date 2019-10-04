package fund

type Fund struct {
	balance float64
}

// NewFund is a basic constructor for Fund struct
func NewFund(i float64) *Fund {
	return &Fund{
		balance: i,
	}
}

// Balance returns the current balance of a Fund
func (f *Fund) Balance() float64 {
	return f.balance
}

// Withdraws
func (f *Fund) Withdraw(i float64) {
	f.balance -= i
}
