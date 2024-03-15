package main

import (
	"fmt"
	"math/rand"
)

/* Быстрая сортировка: из массива выбирается число value, относительно которого будут сравниваться остальные числа.
Массив делится на 2 (числа больше и равные value; числа меньше value).
Для каждого подмассива вызывается рекурсивно эта же функция сравнения, пока длинна массива не станет равна 1.
Отсортированные массивы сливаются в один результирующий массив.*/

func main() {
	var arr []int
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	arr = quickSort(arr)
	fmt.Println(arr)
}

func quickSort(mas []int) []int {
	if len(mas) < 2 {
		return mas
	}

	value := mas[0]
	var great, less []int
	for _, val := range mas[1:] {
		if val >= value {
			great = append(great, val)
			quickSort(great)
		} else {
			less = append(less, val)
			quickSort(less)
		}
	}
	result := append(quickSort(less), value)
	result = append(result, quickSort(great)...)

	return result
}
