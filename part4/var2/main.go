package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/* Остановить горутины можно через каналы или контексты, потому что они выполняются конкурентно.
Контексты предоставляют механизм отмены, который позволяет завершить работу горутины по сигналу отмены контекста. */

func main() {

	const workers = 5
	intCh := make(chan int, 1000)

	ctx, cancelFunc := context.WithCancel(context.Background()) // инициализация контекста и функции cancelFunc

	worker := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case data := <-intCh:
				fmt.Printf("Worker %d received %d\n", id, data)
			case <-ctx.Done(): // сигнал о завершении работы (chan struct{})
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
		i := 0
		for {
			select {
			case intCh <- i:
				i++
			case <-ctx.Done(): // сигнал о завершении работы (chan struct{})
				fmt.Println("Close data chan")
				close(intCh)
				return
			}
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT)

	<-sign

	/* После отмены контекста, вызов ctx.Done() возвращает закрытый канал, и любые операции,
	зависящие от этого контекста, завершаются. */
	cancelFunc()
	fmt.Println("Closed done")

	fmt.Println("Wait for goroutines")
	wg.Wait()
	fmt.Println("End")
}
