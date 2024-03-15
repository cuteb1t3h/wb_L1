package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/* Горутина остановится когда каналы закроются */

func main() {
	intCh := make(chan int, 1000)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case data := <-intCh:
				fmt.Printf("Worker received %d\n", data)
			case <-done:
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
			case <-done:
				close(intCh)
				return
			}
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGINT)

	<-sign

	close(done)
	fmt.Println("Closed done")
	fmt.Println("End")
}
