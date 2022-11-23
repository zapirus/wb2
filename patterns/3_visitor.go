package patterns

import (
	"fmt"
	"math"
)

type shape interface {
	accept(visitor)
}

// Тип объекта 1
type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

// Тип объекта 2
type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

// Тип объекта 3
type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

// Интерфейс visitor с методами для каждого типа объекта
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

// Имплементация интерфейса visitor
type areaCalculator struct {
	area float64
}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = float64(s.side * s.side)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	a.area = float64(math.Pi) * float64(s.radius*s.radius)
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	a.area = float64(s.b * s.l)
}

func ExampleVisitor() {
	fmt.Println("Visitor example")
	fmt.Println()

	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	fmt.Printf("square area is %f\n", areaCalculator.area)
	circle.accept(areaCalculator)
	fmt.Printf("circle area is %f\n", areaCalculator.area)
	rectangle.accept(areaCalculator)
	fmt.Printf("rectangle area is %f\n", areaCalculator.area)
}
