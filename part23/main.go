package main

import "fmt"

/* Для удаления элемента из слайса, используется функция append, которая
добавляет элементы, путем переприсваивания их в тот же слайс, с оператором среза*/

func main() {
	var nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)
	fmt.Println(len(nums))
	i := 4
	nums = Remove(nums, i)
	fmt.Println(nums)
	fmt.Println(len(nums))
}

func Remove(nums []int, i int) []int {
	if i < len(nums) && i >= 0 {
		nums = append(nums[:i], nums[i+1:]...)
	}
	return nums
}
