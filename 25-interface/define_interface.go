package main

type seeder interface {
	seed(int64)
}
type randomizer interface {
	seeder
	pick([]interface{}) interface{}
}

var dice = [...]uint{1, 2, 3, 4, 5, 6}

func rollDice(r randomizer) uint {
	d := make([]interface{}, len(dice))
	for i, n := range dice {
		d[i] = n
	}

	return r.pick([]interface{}(d)).(uint)
}
