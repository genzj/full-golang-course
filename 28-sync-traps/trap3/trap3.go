package trap3

// T for data
type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

// Trap run code
func Trap() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}

// https://play.golang.org/p/LzL4fgKQXA8
