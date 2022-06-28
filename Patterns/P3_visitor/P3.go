/*
Паттерн посетитель позволяет добавлять классам новые методы, не изменяя их, за счет создания нового
класса посетитель, который будет определять необходимые для этих классов методы.
Применимость:
 -  особенно удобен, если необходимо выполнять действия над объектами,
  принадлежащими классу со стабильной структурой.
 - когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
  которые удобно будет вынести в отдельный класс посетитель.
Плюсы:
 - упрощает добавление операций, работающих со сложными структурами объектов
 -  объединяет родственные операции в одном классе.
Минусы:
 - может привести к нарушению инкапсуляции
 - не оправдан если иерархия элементов часто меняется
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct {
	radius float64
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type triangle struct {
	side float64
}

func (t *triangle) accept(v visitor) {
	v.visitForTriangle(t)
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForTriangle(*triangle)
}

type areaCalculator struct {
	area int
}

func (ar *areaCalculator) visitForSquare(s *square) {
	fmt.Printf("Area of square is %d\n", s.side*s.side)
}

func (ar *areaCalculator) visitForCircle(c *circle) {
	fmt.Printf("Area of circle is %0.2f\n", c.radius*c.radius*math.Pi)
}

func (ar *areaCalculator) visitForTriangle(t *triangle) {
	fmt.Printf("Area of triangle is %0.2f\n", math.Sqrt(3)/4*t.side*t.side)
}

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}
	triangle := &triangle{side: 4}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	triangle.accept(areaCalculator)
}
