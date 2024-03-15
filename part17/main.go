package main

import "fmt"

/* Бинарный поиск применяется только к отсортированным массивам. Суть заключается в рекурсивном делении массива
на половины до тех пор, пока не будет найден ключ. Ключ сравнивается со значением середины в полученной половине:
если значение больше, то обрабатывается правая часть массива, если меньше - левая. */

func main() {
	mas := [20]int{9, 13, 16, 32, 42, 45, 48, 63, 63, 64, 66, 67, 70, 71, 74, 74, 77, 78, 87, 96}
	key := 66
	copyMas := make([]int, 20)
	copy(copyMas, mas[:20])
	x := binarySearch(copyMas, key)
	fmt.Println(x)
}

func binarySearch(arr []int, key int) int {
	less := 0
	greater := len(arr) - 1

	for less <= greater {
		mid := (less + greater) >> 1
		midVal := arr[mid]

		if midVal < key {
			less = mid + 1
		} else if midVal > key {
			greater = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
