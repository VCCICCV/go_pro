package main

import "fmt"

//
//import "Test/note"
//
//func main() {
//	//note.Test()
//	//note.IfElse()
//	//note.SwitchCase()
//
//	//a, b := note.Swap("12", "34")
//	//fmt.Print(a, b)
//
//	//note.Array()
//
//	//note.Pointer()
//
//	//note.Type_struct()
//	//note.Test_2()
//
//	//note.Array_2()
//
//	note.Interface()
//}

// interface
//package main
//import (
//	"fmt"
//	"time"
//)
//type Phone interface {
//	call()
//}
//type NokiaPhone struct {
//}
//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}
//
//type IPhone struct {
//}
//
//func (iPhone IPhone) call() {
//	fmt.Println("I am iPhone, I can call you!")
//}
//func main() {
//	var phone Phone
//	phone = new(NokiaPhone)
//	phone.call()
//	phone = new(IPhone)
//	phone.call()
//}

//// goroutine
//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func say(s string) {
//	for i := 0; i < 3; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}
//
//func main() {
//	go say("world")
//	say("hello")·
//}

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")

}
