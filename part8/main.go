package main

import "fmt"

func main() {
	// инициализация переменных
	var num1 int64 = 25
	var num2 int64 = 176
	var i1 uint = 2
	var i2 uint = 5

	fmt.Printf("Set %d bit to 1\n", i1)
	fmt.Printf("Before: %b\t", num1)
	num1 |= 1 << i1 // операция побитового ИЛИ
	fmt.Printf("After: %b\n", num1)

	fmt.Printf("Set %d bit to 0\n", i2)
	fmt.Printf("Before: %b\t", num2)
	num2 &^= 1 << i2 // операция побитового И НЕ
	fmt.Printf("After: %b\n", num2)
}
