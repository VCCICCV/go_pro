package main

import "fmt"

// PROJECT_NAME:Test
// DATE:2022/10/3 12:56
// USER:21005
func main() {
	dog := Dog{
		Animal{name: "dog", age: 21},
		"黑色",
	}
	cat := Cat{
		Animal{name: "cat", age: 21},
		"123",
	}

	dog.eat()
	dog.sleep()
	cat.eat()
	cat.sleep()
}

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat***")

}
func (a Animal) sleep() {
	fmt.Println("sleep****")

}

type Dog struct {
	Animal // 继承了Animal
	color  string
}
type Cat struct {
	Animal // 继承了Animal
	bbb    string
}
