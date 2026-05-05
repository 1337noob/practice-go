package main

import (
	"fmt"
	"sync"
)

func fanin(chans ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for _, ch := range chans {
			wg.Add(1)

			go func() {
				defer wg.Done()

				for v := range ch {
					out <- v
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	result := fanin(ch1, ch2)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for v := range result {
		fmt.Println(v)
	}
}
