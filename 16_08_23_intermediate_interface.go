package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(c Vehicle) {
	c.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}
