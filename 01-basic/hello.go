// 声明当前文件归属的包
package main

// 导入当前文件需要使用的其他包
import (
	"fmt"

	"github.com/genzj/full-golang-course/01-basic/lib"
	"github.com/genzj/full-golang-course/01-basic/util"
)

// 函数
func main() {
	fmt.Printf("Hello, %s and %s\n", globalVar1, constValue1)
	fmt.Printf("Use module lib: %s\n", lib.ExportedConst)
	// 不能使用导入库中的内部函数
	// lib.localVariable()

	fmt.Printf("Use module util: %s\n", util.ExportedConst)
	fmt.Printf("Use escaped var: %s\n", *lib.GetStringPointer())
}
