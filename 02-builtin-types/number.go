package main

var _ uint8  // the set of all unsigned  8-bit integers (0 to 255)
var _ uint16 // the set of all unsigned 16-bit integers (0 to 65535)
var _ uint32 // the set of all unsigned 32-bit integers (0 to 4294967295)
var _ uint64 // the set of all unsigned 64-bit integers (0 to 18446744073709551615)

var _ int8  // the set of all signed  8-bit integers (-128 to 127)
var _ int16 // the set of all signed 16-bit integers (-32768 to 32767)
var _ int32 // the set of all signed 32-bit integers (-2147483648 to 2147483647)
var _ int64 // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

var _ float32 // the set of all IEEE-754 32-bit floating-point numbers
var _ float64 // the set of all IEEE-754 64-bit floating-point numbers

var _ complex64  // the set of all complex numbers with float32 real and imaginary parts
var _ complex128 // the set of all complex numbers with float64 real and imaginary parts

var _ byte // alias for uint8
var _ rune // alias for int32

var x1 int64 = 64

// var x2 int = x1
