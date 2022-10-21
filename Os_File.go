package main

import (
	"fmt"
	"os"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/8 19:35
// USER:21005
// 打开关闭文件
func openCloseFile() {
	// 只读
	f, err := os.Open("go.mod")
	fmt.Printf("openCloseFile.Name:\n%v\n", f.Name())
	// 根据第二格参数，可以读写或者创建
	f2, err2 := os.OpenFile("Rename.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name:%v\n", f2.Name())
	time.Sleep(time.Second * 2)
	err = f.Close()
	fmt.Printf("err:%v\n", err)
	time.Sleep(time.Second * 2)
	err2 = f2.Close()
	fmt.Printf("err:%v\n", err2)
}

// 创建文件
func createFile() {
	f, err3 := os.Create("test.txt")
	fmt.Printf("F.Name:%v\n", f.Name())
	fmt.Printf("err:%v\n", err3)
}
func main() {
	openCloseFile()
}
