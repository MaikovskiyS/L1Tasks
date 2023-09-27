package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	wg := &sync.WaitGroup{}
	numbers := []int{2, 4, 6, 8, 10}
	result := 0.0
	for i := 0; i <= len(numbers)-1; i++ {
		number := numbers[i]
		wg.Add(1)
		go func(int, float64) {
			defer wg.Done()
			res := math.Pow(float64(number), 2)
			result += res
		}(number, result)
		wg.Wait()
	}

	fmt.Printf("result: %v\n", result)

}
