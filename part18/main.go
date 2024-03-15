package main

import (
	"fmt"
	"sync"
)

/* Для безопасного доступа к счетчику структуры используется мьютекс. Он блокирует поток на данный момент времени,
чтобы он был доступен только для одной горутины. После наращивания поток разблокируется для остальных горутин.*/

type Counter struct {
	mx    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	defer c.mx.Unlock()
	c.mx.Lock()
	c.count++
}

func (c *Counter) GetCount() int {
	defer c.mx.Unlock()
	c.mx.Lock()
	return c.count
}

func main() {
	counter := Counter{}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetCount())
}
