package stats

import "github.com/adheeeem/bank.git/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	length := len(payments)
	sum := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
	}
	sum /= types.Money(length)
	return sum
}
