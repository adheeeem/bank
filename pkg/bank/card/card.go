package card

import "github.com/adheeeem/bank/pkg/bank/types"

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active && amount <= card.Balance && amount > 0 && amount <= 20_000_00 {
		card.Balance -= amount
	}
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > 50_000 {
		return
	}
	card.Balance += amount
}
func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			total += card.Balance
		}
	}
	return total
}
func PaymentSources(cards []types.Card) (sources []types.PaymentSource) {
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			cd := types.PaymentSource{
				Type:    card.Name,
				Number:  string(card.PAN),
				Balance: card.Balance,
			}
			sources = append(sources, cd)
		}
	}
	return
}
