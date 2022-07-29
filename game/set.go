package game

type Set struct {
	card    Card
	count   int
	jesters int
}

func (s Set) Rank() int {
	return s.card.Rank()
}

func (s Set) CanBeFollowedBy(other Set) bool {
	return s.count == other.count && s.card.Rank() > other.card.Rank()
}
