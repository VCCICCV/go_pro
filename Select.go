package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/5 20:09
// USER:21005
//var chanInt = make(chan int, 0)
//var chanStr = make(chan string)

func main() {
	//go func() {
	//	chanInt <- 100
	//	chanStr <- "hello world"
	//	defer close(chanInt) //关闭通道
	//	defer close(chanStr)
	//}()
	//for {
	//	select {
	//	case r := <-chanInt:
	//		fmt.Printf("chanInt:%v\n", r)
	//	case r := <-chanStr:
	//		fmt.Printf("chanStr:%v\n", r)
	//	default:
	//		fmt.Printf("default")
	//	}
	//	time.Sleep(1 * time.Nanosecond)
	//}

	//Timer
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now():\n%v\n", time.Now())
	t1 := <-timer.C
	fmt.Printf("t1:\n%v\n", t1)
}
