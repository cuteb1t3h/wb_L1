package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{} // указатель на структуру WaitGroup, которая отвечает за синхронизацию группы горутин
	wg.Add(5)               // определяем значение внутреннего счетчика активных элементов

	var num = [5]int{2, 4, 6, 8, 10} // инициализация массива

	for _, value := range num {
		go sqr(value, wg) /* запуск горутин. В параметры функции передается значение,
		которое будет возводиться в квадрат и указатель на структуру WaitGroup */
	}
	wg.Wait() // ждет завершения всех горутин, блокируя выполнение текущего потока
}

func sqr(value int, wg *sync.WaitGroup) {
	/* defer откладывает выполнение метода Done до конца функции.
	Done нужен для того, чтобы "отмечать" текущую горутину законченной для WG */
	defer wg.Done()
	sqrtValue := value * value
	fmt.Println(sqrtValue)
}
