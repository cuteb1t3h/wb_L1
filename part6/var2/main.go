package main

import (
	"context"
	"fmt"
	"time"
)

/* Горутина остановится функцией отмены по истечению времени жизни контекста */

func main() {
	const timeOut = 5 * time.Second
	intCh := make(chan int, 1000)

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeOut)
	defer cancelFunc()

	go func() {
		for {
			select {
			case data := <-intCh:
				fmt.Printf("Worker received %d\n", data)
			case <-ctx.Done():
				fmt.Println("Done worker")
				return
			}
		}
	}()

	go func() {
		i := 0
		for {
			select {
			case intCh <- i:
				i++
			case <-ctx.Done():
				close(intCh)
				return
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("End")
}
