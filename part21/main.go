package main

import "fmt"

/* Мы имеем интерфейс карты памяти.
Имеем объект ноутбук, без разьема для карты памяти.
Нам нужен адаптер - кардридер,у которого есть usb порт.*/

type Card interface {
	InsertCardIntoPort()
}

type Laptop struct {
}

func (l *Laptop) InsertCardReader() {
	fmt.Println("Card reader inserted")
}

type Phone struct {
}

// InsertCardIntoPort /* Телефон напрямую реализует метод интерфейса*/
func (p *Phone) InsertCardIntoPort() {
	fmt.Println("Card inserted into phone")
}

type CardReader struct {
	lap *Laptop
}

// InsertCardIntoPort /*Адаптер так же напрямую может реализовать метод интерфейса и в нем
// вызвать метод структуры Laptop*/
func (r *CardReader) InsertCardIntoPort() {
	r.lap.InsertCardReader()
	fmt.Println("Connect to laptop via USB")
}

func main() {
	var phone Card = &Phone{} // телефон реализует интерфейс

	phone.InsertCardIntoPort() // вызов метода напрямую к телефону

	laptop := &Laptop{}
	var adapter Card = &CardReader{ // адаптер реализует интерфейс
		lap: laptop,
	}

	adapter.InsertCardIntoPort() // вызов метода ноутбука через адаптер
}
