package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "king", "queen", "jack":
		value = 10
	case "other":
		value = 0
	default:
		value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string
	var card1Value = ParseCard(card1)
	var card2Value = ParseCard(card2)
	var dealerCardValue = ParseCard(dealerCard)
	var sumUpValue = card1Value + card2Value

	switch {
	case card1Value == 11 && card2Value == 11:
		decision = "P"
	case sumUpValue == 21 && (dealerCardValue == 10 || dealerCardValue == 11):
		decision = "S"
	case sumUpValue == 21 && !(dealerCardValue == 10 || dealerCardValue == 11):
		decision = "W"
	case 17 <= sumUpValue && sumUpValue <= 20:
		decision = "S"
	case 12 <= sumUpValue && sumUpValue <= 16 && dealerCardValue <= 6:
		decision = "S"
	case 12 <= sumUpValue && sumUpValue <= 16 && 7 <= dealerCardValue:
		decision = "H"
	case sumUpValue <= 11:
		decision = "H"
	}

	return decision
}
