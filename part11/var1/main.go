package main

import "fmt"

/* Использование 2-х массивов и 1 слайса для реализации. Добавление общих значений из массивов в слайс
происходит во вложенном цикле for. */

func main() {
	first := [7]int{2, 65, 1, 3, 7, 12, 34}
	second := [7]int{25, 9, 1, 34, 0, 2, 7}
	var set []int
	i := 0
	for _, valF := range first {
		for _, valS := range second {
			if valF == valS {
				set = append(set, valF)
				i++
			}
		}
	}
	fmt.Println(set)
}
