package main

import "fmt"

// PROJECT_NAME:Test
// DATE:2022/10/4 0:24
// USER:21005
type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age不能小于0")
	}
	return &Person{name: name, age: age}, nil
}
func main() {
	fmt.Printf("Hello World\n")
	per, err := NewPerson("XiaoWang", 21)
	if err == nil {
		fmt.Printf("per: %v\n", per)
	} else {
		fmt.Printf("err:%v\n", err)
	}
}
