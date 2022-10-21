package main

import (
	"fmt"
	"math/rand"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/4 10:51
// USER:21005

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	time.Sleep(time.Second * 2)
	values <- value
}
func main() {
	defer close(values)
	go send()
	fmt.Printf("wait....\n")
	value := <-values
	fmt.Printf("send:%v\n", value)
	fmt.Printf("end...\n")
}
