package main

import "fmt"

func toCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) / 1.8
}

func main() {
	fahrenheitValues := [...]float32{0, 10, 72, 90, 100}
	celsiusValues := make([]float32, 0, len(fahrenheitValues))

	for _, f := range fahrenheitValues {
		// TODO append the return value of toCelsius(f)
		// to celsiusValues correctly
	}

	fmt.Println("Fahrenheit", fahrenheitValues)
	fmt.Println("Celsius   ", celsiusValues)
}
