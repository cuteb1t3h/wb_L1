package main

import (
	"fmt"
)

func main() {
	tempNums := [8]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := make(map[int][]float64)

	step := 10.0
	/* Каждое значение делится на шаг=10 для того, чтобы определить группу.
	Значение с определенной группой записывается в мапу, где ключ - это группа, а значение - температура. */
	for _, value := range tempNums {
		k := int(value/step) * int(step)
		tempMap[k] = append(tempMap[k], value)
	}
	fmt.Println(tempMap)
}
