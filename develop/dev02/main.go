package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for i := 0; i <= len(numbers)-1; i++ {
		wg.Add(1)
		number := numbers[i]
		go func(int) {
			defer wg.Done()
			res := math.Pow(float64(number), 2)
			fmt.Println(res)
		}(number)

	}
	wg.Wait()
}
