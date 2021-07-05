package shape

import "fmt"

type Shape interface {
	getArea() float64
	getVolume() float64
}

func GetValues(s Shape) {
	s.getArea()
	s.getVolume()
}

type Circle struct {
	Radius float64
}

func (c Circle) getArea() float64 {
	fmt.Println("area of circle")
	return 0
}

func (c Circle) getVolume() float64 {
	fmt.Println("volume of circle")
	return 0
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r *rectangle) getArea() float64 {
	fmt.Println("area of rectangle")
	return 0
}

func (r *rectangle) getVolume() float64 {
	fmt.Println("volume of rectangle")
	return 0
}
