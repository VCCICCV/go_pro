package main

import "fmt"

// PROJECT_NAME:Test
// DATE:2022/10/3 11:48
// USER:21005
func main() {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}

type Flyer interface {
	fly()
}
type Swimmer interface {
	swim()
}
type FlyFish interface {
	Flyer
	Swimmer
}
type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly")
}
func (fish Fish) swim() {
	fmt.Println("swim")
}
