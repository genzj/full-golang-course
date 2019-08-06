package trap2

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

// Trap run code
func Trap() {
	go setup()
	for !done {

	}
	print(a)
}

// https://play.golang.org/p/VXO9X79vI4q
