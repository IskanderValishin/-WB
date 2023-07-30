package main

import "fmt"

type Human struct { // Создание струкртуры Human
	Name string
	Age  int
}
type Action struct { // Создание струкртуры Action
	Human //Встраивание струкрур и методов Human
}

func (n Human) ageString() { // Создание метода Human выводит имя и возраст
	fmt.Printf("Имя человека: %s, Возраст человека: %d \n", n.Name, n.Age)
}
func worked() {
	work := Action{Human{Name: "Vasya", Age: 36}}
	fmt.Printf("%T  %#v \n", work, work)
	fmt.Println(work.Human)
	fmt.Println(work.Age)
	work.ageString()
}

func main() {
	Iskander := Human{}
	Iskander.Name = "Iskander"
	Iskander.Age = 23

	Iskander.ageString()
	worked()
}
