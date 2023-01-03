package stats

import (
	"fmt"
	"github.com/adheeeem/bank.git/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   50_00,
			Category: "korti milli",
		},
		{
			ID:       2,
			Amount:   10_000,
			Category: "visa",
		},
		{
			ID:       3,
			Amount:   35_00,
			Category: "visa",
		},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 6166
}
