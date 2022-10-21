package main

import (
	"log"
	"os"
)
// PROJECT_NAME:Test
// DATE:2022/10/10 1:03
// USER:21005

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("My log:")
	f, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	defer f.Close()
	if err != nil {
		log.Fatal("日志文件错误！")
	}
	log.SetOutput(f)
}
func main() {
	log.Print("my log")
}
