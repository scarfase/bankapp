package card

import "bank/pkg/bank/types"

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(card []types.Card) types.Money {
	var sum types.Money
	for _, card := range card {
		if card.Balance < 0 {
			continue
		}

		if !card.Active {
			continue
		}
		sum += card.Balance
	}
	return sum
}
