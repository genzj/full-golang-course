package main

import "fmt"

// TODO: define two global const containing '32' and '1.8' respectively.
// replace the usages of immediate numbers below with const you just defined

func toCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) / 1.8
}

func toFahrenheit(celsius float32) float32 {
	return 32 + celsius*1.8
}

func main() {
	fmt.Printf("%f C = %f F\n", 13.5, toFahrenheit(13.5))
	fmt.Printf("%f F = %f C\n", 72.0, toCelsius(72.0))
}
