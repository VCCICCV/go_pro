package main

import (
	"fmt"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/5 20:36
// USER:21005
func main() {
	//Timer只执行一次，Ticker可以周期执行
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker 1")
		counter++
		if counter >= 5 {
			break
		}
	}
	ticker.Stop()
}
