package main

import "fmt"

/*конвейер означает, что процессы записи данных в каналы будут происходить параллельно */

func main() {
	intXCh := make(chan int, 6)
	intX2Ch := make(chan int, 6)

	nums := [6]int{2, 56, 42, 8, -31, 12}

	/* для того, чтобы запись и *2 проходили параллельно, создается 2 горутины с записью данных в каналы */
	go func() {
		for _, x := range nums {
			intXCh <- x
		}
		close(intXCh)
	}()

	go func() {
		for x := range intXCh {
			intX2Ch <- x * 2
		}
		close(intX2Ch)
	}()

	for x := range intX2Ch {
		fmt.Println(x)
	}

}
