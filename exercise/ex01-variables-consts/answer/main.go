package main

import "fmt"

const (
	offset   = 32
	multiple = 1.8
)

func toCelsius(fahrenheit float32) float32 {
	return (fahrenheit - offset) / multiple
}

func toFahrenheit(celsius float32) float32 {
	return offset + celsius*multiple
}

func main() {
	fmt.Printf("%f C = %f F\n", 13.5, toFahrenheit(13.5))
	fmt.Printf("%f F = %f C\n", 72.0, toCelsius(72.0))
}
