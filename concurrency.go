package main

import (
	"fmt"
	"sync"
)

func process(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)

	go func() {
		defer wg.Done()
		process(ch)
	}()
	go consume(ch, &wg)
	wg.Wait()
}
