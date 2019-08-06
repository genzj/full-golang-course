package trap1

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

// Trap run code
func Trap() {
	go f()
	g()
}
