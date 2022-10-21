package main

import (
	"fmt"
	"runtime"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/4 11:24
// USER:21005
//
//	func show(s string) {
//		for i := 0; i < 4; i++ {
//			fmt.Println(s)
//		}
//	}
func a() {
	for i := 0; i < 100; i++ {
		fmt.Println("A:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func b() {
	for i := 0; i < 100; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	// Gosched()
	//go show("java")
	//// 主协程
	//for i := 0; i < 4; i++ {
	//	// 切一下，再次分配任务
	//	runtime.Gosched()
	//	fmt.Println("golang")
	//}

	//GOMAXPROCS()
	fmt.Printf("runtime.NumCPU():%v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1) // 修改为1看效果
	go a()
	go b()
	time.Sleep(time.Second)
}
