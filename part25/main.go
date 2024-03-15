package main

import (
	"fmt"
	"time"
)

/* Функция sleep использует канал time.After. Он отправляет значение после заданного времени, с помщью пакета time. */

func main() {
	i := 2
	s := time.Now()

	fmt.Println("Start")
	sleep(i)
	fmt.Println(i, "sec ago")

	e := time.Now()
	d := e.Sub(s)
	fmt.Printf("Время выполнения программы: %v\n", d)
}

func sleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}
