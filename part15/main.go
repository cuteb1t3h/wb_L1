package main

/* Так как слайс - это сслыка на изначальный массив, то для безопасного изменения последнего
нужно скопировать его и после этого делать слайс из нового массива.
Так же justString - это глобальная переменная, поэтому при присваивании слайса сразу туда, в памяти будет храниться
очень большая строка. А при копировании только нужного отрезка, в памяти не сохраняется изначальная огромная строка. */

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	copyBuf := make([]byte, 100)
	copy(copyBuf, v[:100])
	justString = string(copyBuf)
}

/* функция чтоб не горело красным*/
func createHugeString(i int) string {
	return justString
}

func main() {
	someFunc()
}
