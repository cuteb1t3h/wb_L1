package main

import (
	"fmt"
)

/* Для обработки строк, можно использовать руны. Руны представляют собой коды Unicode.
Исходная строка разбивается на отдельные руны, меняется порядок символов и после всего
срез рун преобразуется в строку. */

func main() {
	str := "стажировка😀"
	str = revers(str)
	fmt.Println(str)
}

func revers(str string) string {
	r := []rune(str)

	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	return string(r)
}
