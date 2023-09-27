package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
func main() {
	person := Human{
		name: "john",
	}
	act := Action{
		person: person,
	}
	fmt.Println(act.person.Jump())
}

type Human struct {
	name string
}

func (h *Human) Jump() string {
	return fmt.Sprintf(h.name + "jumping")
}

type Action struct {
	person Human
}
