package card

import (
	"fmt"
	"github.com/adheeeem/bank/pkg/bank/types"
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
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Name:    "card",
			PAN:     "5058 xxxx xxxx 9999",
			Active:  true,
			Balance: 10_000,
		},
		{
			Name:    "another_card",
			PAN:     "5058 xxxx xxxx 2998",
			Active:  true,
			Balance: 50_550,
		},
		{
			Name:    "visa",
			PAN:     "4058 xxxx xxxx 5555",
			Active:  true,
			Balance: 0,
		},
	}
	result := PaymentSources(cards)
	for _, source := range result {
		fmt.Println(source.Number)
	}
	// Output:
	// 5058 xxxx xxxx 9999
	// 5058 xxxx xxxx 2998
}
