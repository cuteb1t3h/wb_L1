package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "хочу на стажировку"
	str = rev(str)
	fmt.Println(str)
}

func rev(str string) string {
	words := strings.Fields(str)
	w := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		w = append(w, words[i])
	}
	return strings.Join(w, " ")
}
