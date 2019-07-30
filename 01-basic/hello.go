// 声明当前文件归属的包
package main

// 导入当前文件需要使用的其他包
import (
	"fmt"

	"./lib"
)

// 全局变量声明和/或定义
var globalVariable string

// 常量声明定义
const name string = "Golang"

// 函数
func main() {
	globalVariable = "Hello"
	fmt.Printf("%s, %s\n", globalVariable, name)
	fmt.Printf("%s\n", lib.ExportedConst)

	fmt.Println(lib.GetStringPointer())
	fmt.Println(lib.GetStringPointer())
	fmt.Println(*lib.GetStringPointer())

	fmt.Printf("%v, %v\n", bit0, mask0)
	fmt.Printf("%v, %v\n", bit1, mask1)
	fmt.Printf("%v, %v\n", bit3, mask3)
}
