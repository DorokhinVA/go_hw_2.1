package credit

import "math"

func Calculate(amount int64, period int64, interestRate int64) (monthlyPay, overpayment, total int64) {
	annuityRate := AnnuityRate(period, interestRate)
	//monthlyPay = int64(annuityRate * float64(amount))
	monthlyPay = int64(math.Round(annuityRate * float64(amount) / 100)) * 100
	total = monthlyPay * period
	overpayment = total - amount
	return
}

func AnnuityRate(period int64, interestRate int64) float64 {
	monthlyRate := math.Round(float64(interestRate) / 12 * 100) / 100 / 100
	pow := math.Pow(1+monthlyRate, float64(period))
	return monthlyRate * pow / (pow - 1)
}
