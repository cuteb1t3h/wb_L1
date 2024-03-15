package main

import "fmt"

func main() {
	var a float64 = 23043231
	var b float64 = 2086423
	fmt.Println(mul(a, b))
	fmt.Println(del(a, b))
	fmt.Println(plus(a, b))
	fmt.Println(diff(a, b))
}

func mul(a float64, b float64) float64 {
	return a * b
}

func del(a float64, b float64) float64 {
	return a / b
}

func plus(a float64, b float64) float64 {
	return a + b
}

func diff(a float64, b float64) float64 {
	return a - b
}
