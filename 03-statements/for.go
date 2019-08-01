package main

import (
	"fmt"
)

func useFor() {
	useForWithNumber()
	useForWithCondition()
	useForWithRange()
}

func useForWithNumber() {
	fmt.Println("useForWithNumber")
	for i := 0; i < 10; i++ {
		fmt.Printf("iteration %d\n", i)
	}
	fmt.Println("---------------")
}

func useForWithCondition() {
	i := 0
	fmt.Println("useForWithCondition")
	for i < 10 {
		fmt.Printf("iteration %d\n", i)
		i++
	}
	fmt.Println("---------------")
}

func useForWithRange() {
	fmt.Println("useForWithRange")
	for x, y := range []int{1, 3, 5, 7, 9} {
		fmt.Printf("iteration %d - element %v\n", x, y)
	}
	fmt.Println("---------------")
}

// fab输出斐波那契中小于等于n的所有元素
// 如n为13时，输出1,1,2,3,5,8,13
// 提示：
// 		斐波那契递推公式为p = p_n-1 + p_n-2
//		输出用fmt.Println(p)每次打印一行，每行一个元素即可
func fab(n uint) {
	//
}
