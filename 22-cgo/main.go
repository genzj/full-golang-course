package main

// #include <stdio.h>
//
// int call_c()
// {
//		printf("in c code\n");
//      int ans = callback(12);
//	    return ans;
// }
import "C"
import "fmt"

func main() {
	fmt.Println(int(C.call_c()))
	// Output: 42
}
