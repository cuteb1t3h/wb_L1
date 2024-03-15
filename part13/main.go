package main

import "fmt"

/* a, b = b, a*/

func main() {
	mas := [4]int{3, 9, 12, 2}
	fmt.Printf("Before: %d %d\n", mas[1], mas[3])

	mas[1], mas[3] = mas[3], mas[1]
	fmt.Printf("After: %d %d\n", mas[1], mas[3])
}
