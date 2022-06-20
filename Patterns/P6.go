/*
Паттерн фабричный метод определяет интерфейс для создания объекта, но
оставляет подклассам решение о том, какой класс инстанцировать. Фабричный метод
позволяет классу делегировать инстанцирование подклассам.
Применимость:
 - классу заранее неизвестно объекты каких классов ему нужно создавать
 - класс спроектирован так, чтобы объекты, которые он создает специфицировались подклассами
 - класс делегирует обязанности одному из нескольких вспомогательных подклассов
 Плюсы:
 - изолирует класс от классов продуктов
 - выделяет производство продуктов в отдельное место, упрощая код
 - упрощает добавление новых продуктов
 - реализует принцип открытости/закрытости

  Минусы:
 - может привести к созданию больших параллельных иерархий классов, так как
  для каждого продукта требуется подкласс создатель
*/

package main

import "fmt"

type iVehicle interface {
	setName(string)
	getName() string
	setSpeed(int)
	getSpeed() int
}

type vehicle struct {
	name  string
	speed int
}

func (v *vehicle) setName(name string) {
	v.name = name
}

func (v *vehicle) getName() string {
	return v.name
}

func (v *vehicle) setSpeed(speed int) {
	v.speed = speed
}

func (v *vehicle) getSpeed() int {
	return v.speed
}

type car struct {
	vehicle
}

func newCar(name string, speed int) iVehicle {
	return &car{
		vehicle{
			name:  name + " car",
			speed: speed,
		},
	}
}

type ship struct {
	vehicle
}

func newShip(name string, speed int) iVehicle {
	return &car{
		vehicle{
			name:  name + " ship",
			speed: speed,
		},
	}
}

func createVehicle(vecType, name string, speed int) iVehicle {
	if vecType == "car" {
		return newCar(name, speed)
	}

	if vecType == "ship" {
		return newShip(name, speed)
	}

	return nil
}

func main() {
	volvo := createVehicle("car", "Volvo", 90)
	titanic := createVehicle("ship", "Titanic", 30)

	fmt.Printf("Given the transport %s with speed %d\n", volvo.getName(), volvo.getSpeed())
	fmt.Printf("Given the transport %s with speed %d\n", titanic.getName(), titanic.getSpeed())
}
