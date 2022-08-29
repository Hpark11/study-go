package main

import (
	"fmt"
	"sync"
)

var data int
var memoryAccess sync.Mutex

func main() {
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is 0.\n")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
