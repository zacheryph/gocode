package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

// Money lets is print money numbers pretty
type Money float64

func (m Money) String() string {
	return fmt.Sprintf("$%.2f", m)
}

var (
	amount   = flag.Float64("amount", 0.0, "loan amount")
	interest = flag.Float64("rate", 0.0, "interest rate")
	months   = flag.Int("months", 360, "length of the loan (360 = 30 years)")
	table    = flag.Bool("table", false, "print amortization table")
	payment  = flag.Float64("payment", 0.0, "override the payment amount")
)

func main() {
	flag.Parse()
	if *amount == 0 || *interest == 0 {
		flag.Usage()
		return
	}

	if *table {
		printAmortizationTable()
	} else {
		periodic := (*interest / 100) / 12
		discount := calculateDiscount(*interest/100, *months)
		payment := *amount / discount
		interestPayment := *amount * periodic

		fmt.Println("Monthly Payment:", Money(payment))
		fmt.Println("First Interest Payment:", Money(interestPayment))
		fmt.Println("First Principal Payment:", Money(payment-interestPayment))
	}
}

func calculateDiscount(interest float64, months int) float64 {
	periodic := interest / 12
	daily := math.Pow(periodic+1, float64(months))
	return (daily - 1) / (periodic * daily)
}

func printAmortizationTable() {
	var totalInterest float64

	periodic := (*interest / 100) / 12
	discount := calculateDiscount(*interest/100, *months)
	balance := *amount
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	monthlyPayment := *payment
	if monthlyPayment == 0 {
		monthlyPayment = balance / discount
	}

	headers := "Period\tOpening\tPayment\tInterest\tPrincipal\tEnding\n"
	writer.Write([]byte(headers))

	for period := 1; period <= *months && balance > 0; period++ {
		interestPayment := balance * periodic
		principalPayment := monthlyPayment - interestPayment

		if principalPayment > balance {
			principalPayment = balance
			monthlyPayment = interestPayment + principalPayment
		}

		line := fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%s\n", period,
			Money(balance), Money(monthlyPayment), Money(interestPayment),
			Money(principalPayment), Money(balance-principalPayment))

		totalInterest += interestPayment

		writer.Write([]byte(line))
		balance -= principalPayment
	}

	writer.Flush()

	writer = tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', 0)
	writer.Write([]byte(fmt.Sprintf("Principal\t%s\n", Money(*amount))))
	writer.Write([]byte(fmt.Sprintf("Total Interest\t%s\n", Money(totalInterest))))
	writer.Write([]byte(fmt.Sprintf("Total Payments\t%s\n", Money(*amount+totalInterest))))
	writer.Flush()
}
