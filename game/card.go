package game

type Card int

const (
	Jester      Card = 13
	Peasant     Card = 12
	Stonecutter Card = 11
	Shepherdess Card = 10
	Cook        Card = 9
	Mason       Card = 8
	Seamstress  Card = 7
	Knight      Card = 6
	Abbess      Card = 5
	Baroness    Card = 4
	EarlMarshal Card = 3
	Archbishop  Card = 2
	Dalmuti     Card = 1
)

func AllCards() []Card {
	cards := make([]Card, 0, 13)
	for i := 1; i <= 13; i++ {
		cards = append(cards, Card(i))
	}
	return cards
}

func (c Card) String() string {
	switch c {
	case Jester:
		return "Jester (13)"
	case Peasant:
		return "Peasant (12)"
	case Stonecutter:
		return "Stonecutter (11)"
	case Shepherdess:
		return "Shepherdess (10)"
	case Cook:
		return "Cook (9)"
	case Mason:
		return "Mason (8)"
	case Seamstress:
		return "Seamstress (7)"
	case Knight:
		return "Knight (6)"
	case Abbess:
		return "Abbess (5)"
	case Baroness:
		return "Baroness (4)"
	case EarlMarshal:
		return "Earl Marshal (3)"
	case Archbishop:
		return "Archbishop (2)"
	case Dalmuti:
		return "Dalmuti (1)"

	default:
		return "Unknown Card"
	}
}

func (c Card) Rank() int {
	return int(c)
}

func (c Card) CardsInDeck() int {
	if c == Jester {
		return 2
	}
	return c.Rank()
}
