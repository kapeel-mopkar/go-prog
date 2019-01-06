package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (ta Triangle) Area() (area float64) {
	fmt.Println("\t\tCalculating area of Triangle")
	area = ta.side * ta.side
	return
}

func (ta Triangle) Perimeter() (perimeter float64) {
	fmt.Println("\t\tCalculating perimeter of Triangle")
	perimeter = 3 * ta.side
	return
}

func (ra Rectangle) Area() (area float64) {
	fmt.Println("\t\tCalculating area of Rectangle")
	area = ra.length * ra.breadth
	return
}

func (ra Rectangle) Perimeter() (perimeter float64) {
	fmt.Println("\t\tCalculating perimeter of Rectangle")
	perimeter = (2 * ra.length) + (2 * ra.breadth)
	return
}

type Triangle struct {
	side float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func CalculateArea(shape Shape) float64 {
	fmt.Println("\tCalculating area of Shape - ", reflect.ValueOf(shape).Type())
	return shape.Area()
}

func CalculatePerimeter(shape Shape) float64 {
	fmt.Println("\tCalculating perimeter of Shape - ", reflect.ValueOf(shape).Type())
	return shape.Perimeter()
}

func main() {
	tra1 := Triangle{34.5}
	//fmt.Println("Area of Triangle : ", tra1.Area())

	ra1 := Rectangle{34.5, 24.5}
	//fmt.Println("Area of Rectangle : ", ra1.Area())

	fmt.Println("Area of Shape Triangle : ", CalculateArea(tra1))
	fmt.Println("Area of Shape Rectangle : ", CalculateArea(ra1))

	fmt.Println("Perimeter of Shape Triangle : ", CalculatePerimeter(tra1))
	fmt.Println("Perimeter of Shape Rectangle : ", CalculatePerimeter(ra1))

	// OUTPUT:
	//$ go run simple_interfaces.go
	// 		Calculating area of Shape -  main.Triangle
	// 				Calculating area of Triangle
	// Area of Shape Triangle :  1190.25
	// 		Calculating area of Shape -  main.Rectangle
	// 				Calculating area of Rectangle
	// Area of Shape Rectangle :  845.25
	// 		Calculating perimeter of Shape -  main.Triangle
	// 				Calculating perimeter of Triangle
	// Perimeter of Shape Triangle :  103.5
	// 		Calculating perimeter of Shape -  main.Rectangle
	// 				Calculating perimeter of Rectangle
	// Perimeter of Shape Rectangle :  118

}
