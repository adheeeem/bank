package stats

import (
	"fmt"
	"github.com/adheeeem/bank.git/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 50_00,
		},
		{
			ID:     2,
			Amount: 10_000,
		},
		{
			ID:     3,
			Amount: 35_00,
		},
	}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 6166
}
