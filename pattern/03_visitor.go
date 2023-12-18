package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель — это поведенческий паттерн, который позволяет добавить
новую операцию для целой иерархии классов, не изменяя код этих
классов.

Применение:
    Если имеются различные объекты разных классов с разными интерфейсами,
	но над ними нужно совершать операции, зависящие от конкретных классов
    Если необходимо над структурой выполнить различные, усложняющие структуру операции;
    Если часто добавляются новые операции над структурой.

Плюсы:
    Упрощается добавление новых операций
    Объединение родственных операции
 	Посетитель может запоминать в себе какое-то состояние по мере обхода контейнера.

Минусы:
	Затруднено добавление новых классов,
	поскольку нужно обновлять иерархию посетителя и его сыновей.
*/

type Shape interface {
	accept(Visitor)
}

type Visitor interface {
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

type AreaCalculator struct{}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for a circle")
}
func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for a rectangle")
}

type MiddleCoordinates struct{}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle coordinates for a circle")
}
func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Calculating middle coordinates for a rectangle")
}

func main() {
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	middleCoordinates := &MiddleCoordinates{}
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
