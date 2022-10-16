package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	test19()
}

func test1() {
	// writeStream := make(chan<- interface{})
	// readStream := make(<-chan interface{})
	// <-writeStream
	// readStream <- struct{}{}
}

func test2() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 {
			return
		}
		stringStream <- "Hello Channels!"
	}()
	fmt.Println(<-stringStream)
}

func test3() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello Channels!"
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutation)
}

func test4() {
	valueStream := make(chan interface{})
	close(valueStream)

	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v", ok, integer)
}

func test5() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

func test6() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func test7() {
	// var dataStream chan interface{}
	// dataStream = make(chan interface{}, 4)
}

func test8() {
	// a := make(chan int)
	// b := make(chan int, 0)
	// c := make(chan rune, 4)
}

func test9() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 4; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v. \n", integer)
	}
}

func test10() {
	var dataStream chan interface{}
	<-dataStream
}

func test11() {
	var dataStream chan interface{}
	dataStream <- struct{}{}
}

func test12() {
	var dataStream chan interface{}
	close(dataStream)
}

func test13() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func test14() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select {
	case <-c1:
	case <-c2:
	case c3 <- struct{}{}:
	}
}

func test15() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func test16() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func test17() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}

func test18() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

func test19() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

// func twoWayChannel() {
// 	var dataStream chan interface{}
// 	dataStream = make(chan interface{})
// }

// func receiveOnlyChannel() {
// 	var dataStream <-chan interface{}
// 	dataStream = make(<-chan interface{})
// }

// func sendOnlyChannel() {
// 	var dataStream chan<- interface{}
// 	dataStream = make(chan<- interface{})
// }
