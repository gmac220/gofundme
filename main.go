package main

import (
	"flag"
	"fmt"

	"github.com/gmac220/gofundme/fund"
)

func main() {
	initbalance := flag.Float64("balance", 100, "default balance")
	name := flag.String("name", "Mehrab", "default name")
	flag.Parse()

	fmt.Println(*name)
	fund := fund.NewFund(*initbalance)

	lock := make(chan bool)

	go func() {
		fund.Withdraw(50)
		lock <- true
	}()

	// fmt.Println("Hit enter to continue")
	// fmt.Scanln()

	<-lock
	fund.Withdraw(45)

	fmt.Println(fund.Balance())
}
