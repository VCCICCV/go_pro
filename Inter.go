package main

import "fmt"

// PROJECT_NAME:Test
// DATE:2022/10/3 10:51
// USER:21005
type USB interface {
	read()
	write()
}
type Computer struct {
	name string
}
type Mobile struct {
	model string
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("m.read***")
}
func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("m.write***")
}
func (c Computer) read() {
	fmt.Printf("c.name:%v\n", c.name)
	fmt.Print("read****\n")
}
func (c Computer) write() {
	fmt.Printf("c.name:%v\n", c.name)
	fmt.Print("write****\n")
}
func main() {
	c := Computer{
		name: "小米",
	}
	c.read()
	c.write()
	m := Mobile{
		model: "5G",
	}
	m.read()
	m.write()
}
