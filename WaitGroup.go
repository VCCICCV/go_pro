package main

import (
	"fmt"
	"sync"
)

// PROJECT_NAME:Test
// DATE:2022/10/4 11:13
// USER:21005
var wg1 sync.WaitGroup

func hello(i int) {
	defer wg1.Done()
	fmt.Printf("Hello Goroutine: %d\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go hello(i)
	}
	wg1.Wait()
}
