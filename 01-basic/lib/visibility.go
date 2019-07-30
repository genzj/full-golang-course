package lib

import "fmt"

// 以下所说的名称规则，适用于变量、常量、函数、类型、接口等几乎所有Go对象

// 小写开头的名称只能在当前包内使用
const localConst = "local in package"

// 大写开头的名称能被其他包导入使用
const ExportedConst = "exported to other package"

func localVariable() {
	// 函数体内的局部变量只能在函数及闭包内被访问
	var local1 string
	local1 = "local variable 1"

	// 也可以在声明时赋值，从而利用类型推导
	var local2 = "local variable 2"

	// 声明+赋值可以使用简化操作符:=
	local3 := "local variable 3"

	// 所有的局部变量必须被使用（右值）过至少一次
	// 只定义/赋值而不被使用的变量会被当作错误
	fmt.Printf("%s, %s, %s", local1, local2, local3)
}
