package main

import "fmt"

// 先声明，后在某个函数中赋值
var globalVar1 string

// init是每个文件的初始化（回调）函数，可以在此处初始化各个全局变量
func init() {
	globalVar1 = "global variable 1"
}

// 也可以在声明的同时赋值
var globalVar2 string = "global variable 2"

// 因为有了赋值，因此变量类型可以被推导，因此推荐写作
var globalVar3 = "global variable 3"

// 推荐将逻辑相关的变量放在一起声明
var (
	errIO      = fmt.Errorf("IO error")
	errNetwork = fmt.Errorf("network error")
	errFile    = fmt.Errorf("file error")
)

// 另外一种方法
var errA, errB = fmt.Errorf("A"), fmt.Errorf("B")

// 常量和变量的声明类似，但必须在声明时赋值，赋值必须为常量
const constValue1 = "const value 1"
const constValue2 string = "const value 2"

// 错误：声明时未赋值
// const badConst1 string

// 错误：赋值不是常量
// const badConst2 = fmt.Errorf("error")
