package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	or := make(chan interface{})

	waitForClose := func(ch <-chan interface{}) {
		for val := range ch {
			or <- val
		}
		wg.Done()
	}

	wg.Add(1)
	for _, ch := range channels {
		go waitForClose(ch)
	}

	go func() {
		wg.Wait()
		close(or)
	}()

	return or
}
