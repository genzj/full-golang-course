package main

func addOne(x interface{}) {
	switch i := x.(type) {
	// TODO increase x by one properly by its type then print the result:
	//   int / float32: algebraic addition
	//   rune: next char in the alphabetic table, (a+1=b, b+1=c, ..., z+1=a)
	//   others: print error message
	}
}

func main() {
	addOne(1)
	addOne(100.12)
	addOne('a')
	addOne(true)
}
