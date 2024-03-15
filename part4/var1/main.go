package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	const workers = 5             // выбор кол-ва воркеров
	intCh := make(chan int, 1000) // буфферизированный канал, для передачи данных
	done := make(chan struct{})   // пустой канал для завершения горутин (struct{} 0 байт)

	/* функция worker запускает бесконечный цикл чтения данных из канала.
	Пока в канале intCh есть данные и он не закрыт, в data записываются поступающие значения.
	select выбирает тот case, который быстрее отдаст данные из канала. После закрытия канала done
	он мгновенно вернет нулевое значение и false. Выполнится case <-done и выход из функции.
	*/
	worker := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case data := <-intCh:
				fmt.Printf("Worker %d received %d\n", id, data)
			case <-done:
				fmt.Printf("Done %d worker\n", id)
				return
			}
		}
	}

	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)         // увеличение счетчика горутин перед запуском
		go worker(i, &wg) //запуск воркера
	}

	/* анонимная функция, которая запускает бесконечный цикл записи данных в канал.
	Пока канал done не закрыт, select будет выбирать case записи. После закрытия канала done
	выполнится закрытие канала intCh и выход из функции.
	*/
	go func() {
		i := 0
		for {
			select {
			case intCh <- i:
				i++
			case <-done:
				fmt.Println("Close data chan")
				close(intCh)
				return
			}
		}
	}()

	sign := make(chan os.Signal, 1)                   // канал для получения сигнала о прерывании программы
	signal.Notify(sign, os.Interrupt, syscall.SIGINT) // подписка на сигналы прерывания

	<-sign //ожидание сигнала о прерывании
	/* пока в канал sign не поступят данные, дальнейший код не будет выполняться */

	close(done) // закрытие канала done
	fmt.Println("Closed done")

	fmt.Println("Wait for goroutines")
	fmt.Println("End")
}
