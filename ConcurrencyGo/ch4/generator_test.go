package main

import "testing"

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
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

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for _, v := range values {
			select {
			case <-done:
				return
			case valueStream <- v:
			}
		}
	}()
	return valueStream
}

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {

	}
}

func BenchmarkTyped(b *testing.B) {
	
}
