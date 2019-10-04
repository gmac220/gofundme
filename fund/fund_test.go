package fund

import (
	"log"
	"strconv"
	"testing"
)

func Test_NewFund(t *testing.T) {
	fund := NewFund(100.00)

	if fund.Balance() == 100 {
		log.Printf("Creating fund with 100 works")
	} else {
		t.Errorf("Expected 100 but got: " + strconv.FormatFloat(fund.Balance(), 'f', 2, 64))
	}
}

func Test_Withdraw(t *testing.T) {
	fund := NewFund(100.00)

	fund.Withdraw(100)
	if fund.Balance() == 0 {
		log.Printf("Withdrawing 100 works")
	} else {
		t.Errorf("Expected 100 but got: " + strconv.FormatFloat(fund.Balance(), 'f', 2, 64))
	}
}

func Bench_Fund(b *testing.B) {
	fund := NewFund(float64(b.N))

	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() == 0 {
		log.Println("All funds withdrawn successfully")
	} else {
		b.Errorf("Something went wrong")
	}
}
