package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000,
			Active:  true,
		},
		{
			Balance: 15_000,
			Active:  true,
		},
		{
			Balance: 5_000,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	// Output: 25000
}
