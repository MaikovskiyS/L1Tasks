package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
const elementIndex = 3

func main() {

	//#1 сохраняем порядок элементов; мутируем исходный массив
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	resultV1 := deleteElementV1(arr1, elementIndex)
	fmt.Println(resultV1)

	//#2 Не сохраняем порядок элементов; мутируем исходный массив
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	resultV2 := deleteElementV2(arr2, elementIndex)
	fmt.Println(resultV2)

	//#3 сохраняем порядок элементов; не изменяем исходный массив
	arr3 := []int{1, 2, 3, 4, 5, 6, 7}
	resultV3 := deleteElementV3(arr3, elementIndex)
	fmt.Println(resultV3)

}
func deleteElementV1(arr []int, i int) []int {

	return append(arr[:i], arr[i+1:]...)
}
func deleteElementV2(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
func deleteElementV3(arr []int, i int) []int {
	newArr := make([]int, 0)
	newArr = append(newArr, arr[:i]...)
	return append(newArr, arr[i+1:]...)
}
