package main

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

// 重复的“= iota”可以省略，编译器会自动补上
// 这是声明const时省略初始化的特殊情况
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // 该项不导出
)

// iota一般只在const定义块内用，否则总是0
const x = iota // x == 0
const y = iota // y == 0

// iota随const定义行增加计数，即便该行没有使用iota
const (
	a = 1 << iota // a == 1  (iota == 0)
	b = 1 << iota // b == 2  (iota == 1)
	c = 3         // c == 3  (iota == 2, unused)
	d = 1 << iota // d == 8  (iota == 3)
)

// 换言之，同一行内iota是不变的
const (
	bit0, mask0 = 1 << iota, (1 << iota) - 1 // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                              // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                     //                        (iota == 2, unused)
	bit3, mask3                              // bit3 == 8, mask3 == 7  (iota == 3)
)

const (
	u         = iota * 42 // u == 0     (untyped integer constant)
	v float64 = iota * 42 // v == 42.0  (float64 constant)
	w         = iota * 42 // w == 84    (untyped integer constant)
)
