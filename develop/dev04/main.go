package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
func main() {
	wg := &sync.WaitGroup{}
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	ch := make(chan int, 10)
	fmt.Println("select number of workers:")
	var val int
	fmt.Scan(&val)
	for i := 1; i != val+1; i++ {
		worker := i
		wg.Add(1)
		go func(int, chan int, chan os.Signal) {
			for v := range ch {
				fmt.Println("worker №:", worker)
				fmt.Println("value from ch:", v)
			}
			wg.Done()
			fmt.Println("worker done", worker)
		}(worker, ch, sigch)
	}
	i := 0
	for {
		select {
		case <-sigch:
			fmt.Println("exit by signal")
			close(ch)
			wg.Wait()
			fmt.Println("workers end")
			close(sigch)
			return
		default:
			ch <- i
			i++
		}

	}

}

/*
воркеры завершают свою работу только после закрытия канала с данными.
канал закрывается, воркеры доделывают работу, закрываются, как только последний воркер закрылся, выходим из main
*/
