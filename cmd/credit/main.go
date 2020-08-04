package main

import (
	"fmt"
	"github.com/DorokhinVA/go_hw_2.1/pkg/credit"
)

func main() {
	monthlyPay, overpayment, total := credit.Calculate(1_000_000_00, 36, 20)

	fmt.Println("Monthly Pay: ", monthlyPay)
	fmt.Println("Overpayment: ", overpayment)
	fmt.Println("Total: ", total)
}
