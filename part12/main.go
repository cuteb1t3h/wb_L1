package main

import "fmt"

func main() {
	str := [5]string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	/* Цикл проходит по массиву str и добавляет каждый элемент как ключ в мапу.
	Так как ключ уникален, то каждый уникальный элемент str запишется в множество*/
	for _, item := range str {
		set[item] = true
	}
	fmt.Println(set)
}
