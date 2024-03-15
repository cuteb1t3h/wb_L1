package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // синхронизация горутин

	var nums = [5]int{2, 4, 6, 8, 10}
	sum := 0

	wg.Add(5) //определяем 5 воркеров
	for _, value := range nums {
		go sqrtSum(value, &sum, wg)
	}
	wg.Wait() // ждем выполнения всех воркеров
	fmt.Println(sum)
}

func sqrtSum(value int, sum *int, wg *sync.WaitGroup) {
	defer wg.Done() // вычитается завершенная горутина из кол-ва воркеров
	value = (value) * (value)
	*sum += value
}
