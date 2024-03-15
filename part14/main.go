package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{}

	//val = 7
	//val = "string"
	//val = true
	val = make(chan int)

	valType := reflect.TypeOf(val) // пакет reflect позволяет в процессе программы определить тип переменной

	fmt.Println(valType)
}
