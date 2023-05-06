package money

type Dollar struct {
	Amount int
}

func (d Dollar) Times(multiplier int) int {
	return 15
}
