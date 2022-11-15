package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	test2()
}

func test1() {
	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		wg.Add(len(channels))

		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}
}

func test2() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		repeatStream := make(chan interface{})
		go func() {
			defer close(repeatStream)
			for {
				select {
				case <-done:
					return
				case repeatStream <- fn():
				}
			}
		}()

		return repeatStream
	}

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()

		return intStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()

		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	// start := time.Now()
	rand := func() interface{} { return rand.Intn(5000000) }

	randIntStream := toInt(done, repeatFn(done, rand))
	for i := range randIntStream {
		fmt.Println(i)
	}
}
