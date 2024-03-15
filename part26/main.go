package main

import "fmt"

/* Создается мапа для символов, которые проходятся в строке.
В цикле проверяется встречался ли этот символ ранее, без учета регистра. И если встречался, то функция возвращает false.
Иначе символ добавляется в мапу.*/

func main() {
	strings := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range strings {
		fmt.Printf("Строка \"%s\": %t\n", str, isUnique(str))
	}
}

func isUnique(str string) bool {
	seen := make(map[rune]bool)

	for _, char := range str {
		char = toLower(char)

		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

// Приводим к нижнему регистру (аски)
func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return char + 32
	}
	return char
}
