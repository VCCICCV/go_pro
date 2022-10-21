package main

import (
	"fmt"
	"os"
)

// PROJECT_NAME:Test
// DATE:2022/10/8 23:47
// USER:21005
func write() {
	f, err := os.OpenFile("Rename.txt", os.O_RDWR|os.O_APPEND, 0755) //读写和追加
	if err != nil {
		fmt.Println("Error write")
	} else {
		fmt.Println("Rename.txt successfully")
	}
	f.Write([]byte("jkjdsdksj"))
	f.Close()
}

// 写字符串
func writeWString() {
	f, _ := os.OpenFile("REname.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.WriteString("写字符串") //追加模式
	f.Close()
}
func writeAt() {
	f, _ := os.OpenFile("Rename.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa1234556"), 3) // 覆盖
	f.Close()
}
func main() {
	//write()
	//writeWString()
	writeAt()
}
