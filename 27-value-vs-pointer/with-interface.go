package main

type anyInterface interface {
	fun()
}

func someFunction(i anyInterface) {}

type someStruct struct{}

func (s *someStruct) fun() {}

func useInterface() {
	s := someStruct{}
	// 下面语句是合法的，go会自动取s的地址传给fun
	s.fun()
	// 👇下面调用不能编译通过，因为someStruct没有实现方法fun，不是anyInterface类型
	// someFunction(s)
	// 👇下面调用可以工作，因为方法fun绑定值*someStruct上
	someFunction(&s)

	// go中的自动解引用，自动转指针等语法特性会在此时让人困惑😞
	// 1. T和*T是两种完全不同的类型（强类型），赋值、传参、接口断言时有区别
	// 2. 但调用x.fun()时，无论x是T还是*T均可，不用显式写(&x).fun()
	// 3. go中没有C的指针访问操作符(->)
}
