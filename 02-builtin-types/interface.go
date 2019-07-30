package main

import "fmt"

func printInterface() {
	var x map[string]interface{}
	x = make(map[string]interface{}, 3)

	x["key-1"] = 12345
	x["key-2"] = "abcde"
	x["key-3"] = 3.1415927

	for k, v := range x {
		fmt.Printf("%s - %#v (%T)\n", k, v, v)
	}
}

func typeAssertion() {
	var x map[string]interface{}
	x = make(map[string]interface{}, 3)

	x["key-1"] = 12345
	x["key-2"] = "abcde"
	x["key-3"] = 3.1415927

	// var i int = int(x["key-1"])
	key := "key-2"
	i, ok := x[key].(int)
	if ok {
		fmt.Printf("%#v is %T\n", i, i)
	} else {
		fmt.Printf("%#v isn't of type %T\n", x[key], x[key])
	}

	s, ok := x[key].(string)
	if ok {
		fmt.Printf("%#v is %T\n", s, s)
	} else {
		fmt.Printf("%#v isn't of type %T\n", x[key], x[key])
	}

	// ... ...

	for _, v := range x {
		switch i := v.(type) {
		case int:
			fmt.Printf("%d is an integer\n", i)
		case string:
			fmt.Printf("%s is a string\n", i)
		default:
			fmt.Printf("%#v is of unexpected type", i)
		}
	}

}
