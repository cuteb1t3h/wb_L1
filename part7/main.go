package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	people := make(map[string]int)
	name := [5]string{"Kate", "Nikola", "Anna", "Peter", "Volt"}

	var mutex sync.Mutex // определение мьютекса для безопасного доступа к общей мапе

	const workers = 3
	stringCh := make(chan string) // не буфферизированный канал, для передачи данных

	ctx, cancelFunc := context.WithCancel(context.Background())

	worker := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case data := <-stringCh:
				mutex.Lock() // блокирование доступа к мапе только для одной горутины в данный момент времени
				age := rand.Intn(26) + 5
				people[data] = age
				fmt.Printf("Worker %d received %s and %d\n", id, data, age)
				mutex.Unlock() // разблокировка доступа остальным горутинам после освобождения мапы
			case <-ctx.Done():
				fmt.Printf("Done %d worker\n", id)
				return
			}
		}
	}

	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	go func() {
		for i := 0; i < 5; i++ {
			stringCh <- name[i]
		}
		return
	}()

	cancelFunc()
	wg.Wait()

	fmt.Println("Closed done")
	fmt.Println("End")
}
