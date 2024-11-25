package main

import (
	"fmt"
	"sync"
)

func produce(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int, 5)
	go func() {
		defer wg.Done()
		produce(ch)
	}()
	go consume(ch, &wg)

	wg.Wait()
}
