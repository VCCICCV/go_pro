package main

import "fmt"

// PROJECT_NAME:Test
// DATE:2022/10/5 20:01
// USER:21005
var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()
	//r := <-c
	//fmt.Printf("r:%d\n", r)
	//r = <-c
	//fmt.Printf("r:%d\n", r)
	//r = <-c
	//fmt.Printf("r:%d\n", r)
	for i := 0; i < 20; i++ {
		r := <-c
		fmt.Printf("r:%d\n", r)
	}
}
