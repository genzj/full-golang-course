package main

import "fmt"

func useMap() {
	// map是可以动态扩增的键值对序列

	// 可以通过make创建，以便预分配空间(长度不可指定)
	// m := make(map[string]int, 10)

	// 也可以用简便写法，方便初始化
	m := map[string]int{"three": 3, "four": 4, "zero": 0}
	m["one"] = 1

	// 访问一个不存在的键会得到空值
	fmt.Printf("non-existent key->%#v\n", m["xxx"])

	// 区分空值和不存在的方法：
	target := "zero"
	if val, exist := m[target]; exist {
		fmt.Printf("exist value: %s->%#v\n", target, val)
	} else {
		fmt.Printf("non-exist key: %#v\n", target)
	}

	for key, value := range m {
		// 遍历过程中可以增加项目，但是增加项目遍历与否不确定
		// m[key+"_100"] = value + 100
		// 遍历过程中也可以删除项目
		// delete(m, key)
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	fmt.Printf("final map (len=%d): %#v\n", len(m), m)
}

func trap() {
	type data struct {
		x string
	}
	m := map[string]data{
		"one": data{x: "one"},
	}

	// error
	// m["one"].x = "two"

	{
		tmp := m["one"]
		tmp.x = "two"
		m["one"] = tmp
	}

	fmt.Printf("final value: %#v\n", m["one"])
}

func main() {
	useMap()
	trap()
}
