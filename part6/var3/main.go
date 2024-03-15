package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*Горутина остановится функцией отмены, при нажатии Ctrl+C*/

func main() {
	intCh := make(chan int, 1000)

	ctx, cancelFunc := context.WithCancel(context.Background())

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

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT)

	<-sign

	cancelFunc()
	fmt.Println("Closed done")
	fmt.Println("End")
}
