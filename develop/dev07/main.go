package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/
func main() {
	wg := sync.WaitGroup{}
	m := &sync.RWMutex{}
	mapa := make(map[int]bool)
	for i := 0; i <= 4; i++ {
		worker := 1 + i
		wg.Add(1)
		go func(*sync.RWMutex, map[int]bool, int) {
			defer wg.Done()
			id := 0
			for {
				fmt.Printf("Worker № %+v recording in map\n", worker)
				m.Lock()
				mapa[id] = true
				m.Unlock()
				id++
				if len(mapa) == 50 {
					return
				}
			}
		}(m, mapa, worker)
	}
	wg.Wait()
	fmt.Println(mapa)
	fmt.Println("main end")
}
