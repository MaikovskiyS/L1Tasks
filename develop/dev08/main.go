package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
const bitNumber int64 = 5
const number int64 = 100

func main() {
	changeBit(number, bitNumber)
}
func changeBit(number int64, bit int64) {
	switchToOne := false
	switch switchToOne {
	case true:
		number |= 1 << bit
	default:
		number &^= 1 << bit
	}

	fmt.Println("result is - ", number)
}
