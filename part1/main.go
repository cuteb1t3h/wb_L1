package main

import "fmt"

// создание структуры Human
type Human struct {
	Name string
	Age  int
}

/* Для того, чтобы поля, методы структуры и сама структура были доступны вне текущего пакета (main),
используются названия написанные с заглавной буквы*/

// метод структуры Human
func (h *Human) Hello() {
	fmt.Printf("Hi! My name is %s. I'm %d years old.\n", h.Name, h.Age)
}

// встраивание структуры Human в Action
type Action struct {
	Human
}

/* Реализацию структуры Action можно рассматривать как
аналог наследования в объектно-ориентированном программировании.*/

// метод структуры Action, который использует поле и метод структуры Human
func (a *Action) Swim() {
	a.Hello()
	fmt.Println("And I learned to swim when I was", a.Age-6)
}

/*За счет того, что структура Human "встроена" в структуру Action, поля и методы становятся доступными,
несмотря на то, что они не определены непосредственно в структуре Action.*/

func main() {
	// иницализация экземпляра структуры Action вложенным экземлпяром Human
	person := Action{
		Human{
			Name: "Alice",
			Age:  20,
		},
	}
	// вызов метода Swim для инициализированного объекта
	person.Swim()
}
