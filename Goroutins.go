package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/4 1:12
// USER:21005
func main() {
	go showMsg("Welcome") // 启动了一个协程
	go showMsg("java")
	showMsg("test")
	//time.Sleep(time.Millisecond * 10000)
	fmt.Print("main end")
}
func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%s\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
