package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/4 22:33
// USER:21005
//var i int = 100
//var wg sync.WaitGroup
//var lock sync.Mutex
//func add() {
//	defer wg.Done()
//	lock.Lock()
//	i += 1
//	fmt.Printf("i++:%v\n", i)
//	lock.Unlock()
//}
//func sub() {
//	defer wg.Done()
//	lock.Lock()
//	i -= 1
//	fmt.Printf("i--:%v\n", i)
//	lock.Unlock()
//}
//func main() {
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go add()
//		wg.Add(1)
//		go sub()
//	}
//	wg.Wait()
//	fmt.Printf("end:%v\n", i)
//}

// 原子操作 cas xompare and swap
var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}
func sub() {
	atomic.AddInt32(&i, -1)
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second)
	fmt.Printf("end:%v\n", i)

}
