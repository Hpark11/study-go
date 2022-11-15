package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valStream <- v:
				}
			}
		}
	}()
	return valStream
}

func take(done <-chan interface{}, c <-chan interface{}, n int) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case valStream <- <-c:
			}
		}
	}()

	return valStream
}

func PerformWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}

func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return file
}
