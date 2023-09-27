package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/
const searchElement = 5

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := binarySearch(arr, searchElement)
	fmt.Println(result)
}
func binarySearch(arr []int, i int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		midIndex := (left + right) / 2
		midValue := arr[midIndex]
		if midValue == i {
			return midIndex
		} else if midValue < i {
			left = midIndex + 1
		} else {
			right = midIndex - 1
		}
	}
	return 0
}

/*
1. Начните с определения границ поиска. Установите начальную границу (left) в начало списка и конечную границу (right) в его конец.

2. Найдите средний элемент между начальной и конечной границей.

3. Сравните средний элемент с искомым значением.

4. Если средний элемент совпадает с искомым значением, возвращаем его индекс (или значение, в зависимости от реализации).

5. Если средний элемент больше искомого значения, обновите конечную границу (right) на позицию перед средним элементом и перейдите к шагу 2.

6. Если средний элемент меньше искомого значения, обновите начальную границу (left) на позицию после среднего элемента и перейдите к шагу 2.

7. Повторяйте шаги 2-6 до тех пор, пока не будет найден искомый элемент или пока начальная граница (left) не станет больше конечной границы (right).

Бинарный поиск имеет временную сложность O(log n), где "n" - это количество элементов в списке. Это означает, что в худшем случае алгоритм будет выполняться
за логарифмическое время, что является значительно более эффективным, чем линейный поиск.
*/
