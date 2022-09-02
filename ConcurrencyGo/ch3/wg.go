package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// salutation := "hello"

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	// fmt.Println("Hello")
	// 고루틴 자신이 생성된 곳과  동일한 주소 공간에서 실행되면 프로그램은 고루틴 내부에서 외부변수를 수정한다.
	// 	salutation = "welcome"
	// }()

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)

		go func(saludation string) {
			defer wg.Done()
			fmt.Println(saludation)
		}(salutation)
	}

	wg.Wait()
	// fmt.Println(salutation)
}
