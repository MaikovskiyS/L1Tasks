package main

import "fmt"

/*
1.значение нулевого элемента изменилось, так как слайсы содержат адрес памяти исходного массива.
2. append не видно за пределами функции, так как при добавлении элемента мы привышаем capacity, из за этого создается новый массив.
*/
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
