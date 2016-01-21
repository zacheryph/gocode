package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	amount   = flag.Float64("amount", 0.0, "loan amount")
	interest = flag.Float64("rate", 0.0, "interest rate")
	months   = flag.Int("months", 360, "length of the loan (360 = 30 years)")
)

func main() {
	flag.Parse()
	if *amount == 0 || *interest == 0 {
		flag.Usage()
		return
	}

	discount := calculateDiscount(*interest/100, *months)
	payment := *amount / discount

	fmt.Printf("Monthly Payment: $%.2f\n", payment)

	periodic := (*interest / 100) / 12
	interestPayment := *amount * periodic
	fmt.Printf("First Interest Payment: $%.2f\n", interestPayment)
	fmt.Printf("First Principal Payment: $%.2f\n", payment-interestPayment)

}

func calculateDiscount(interest float64, months int) float64 {
	periodic := interest / 12
	daily := math.Pow(periodic+1, 360)
	return (daily - 1) / (periodic * daily)
}
