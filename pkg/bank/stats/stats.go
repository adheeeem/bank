package stats

import "github.com/adheeeem/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	length := len(payments)
	sum := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
	}
	sum /= types.Money(length)
	return sum
}
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			total += payment.Amount
		}
	}
	return total
}
