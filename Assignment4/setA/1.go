// wap in go language to create an interface shape that includes area and perimeter.
// implements these methods in circle and rectanglee types

package main

import "fmt"

type Shape interface {
	area()

	perimeter()
}

type Circle struct {
	pie    float32
	radius float32
}
type Rectangle struct {
	length float32
	width  float32
}

func (c Circle) area() {
	fmt.Printf("The area of circle is : %v\n", c.pie*c.radius*c.radius)
}
func (c Circle) perimeter() {
	fmt.Printf("The perimeter of circle is : %v\n", 2*c.pie*c.radius)

}
func (r Rectangle) area() {
	fmt.Printf("The area of Rectangle is :%v\n", r.length*r.width)

}
func (r Rectangle) perimeter() {
	fmt.Printf("The perimeter of Rectangle is :%v\n", 2*(r.length+r.width))

}
func main() {
	c := Circle{
		pie:    3.14,
		radius: 5,
	}
	c.area()
	c.perimeter()

	r := Rectangle{
		length: 5,
		width:  3,
	}
	r.area()
	r.perimeter()

}

// Area of Circle = pie * r * r
// Perimeter of circle = 2*Pie*r
// Area of Rectangle = lw
// perimeter of rectangle = 2(l+w)
