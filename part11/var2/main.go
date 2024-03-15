package main

import "fmt"

/* Использование 2-х массивов и 1 мапы для реализации. Сначала в мапу добавляются значения из 1 массива,
а потом итерируется по второму, проверяя наличие в мапе. */

func main() {
	first := [7]int{2, 65, 1, 3, 7, 12, 34}
	second := [7]int{25, 9, 1, 34, 0, 2, 7}
	set := make(map[int]bool)
	for _, val := range first {
		set[val] = true
	}

	var intSec []int
	for _, val := range second {
		if set[val] {
			intSec = append(intSec, val)
		}
	}
	fmt.Println(intSec)
}
