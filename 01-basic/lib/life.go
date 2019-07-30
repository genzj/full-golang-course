package lib

// GetStringPointer returns a ref to string
func GetStringPointer() *string {
	var x = "local string variable"
	// 返回局部变量的指针在C/C++里是大忌
	// go里如此操作是否安全？
	return &x
}
