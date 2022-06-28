/*
Цепочка вызовов/обязвнностей - это поведенческий паттерн, позволяющий
избежать жёсткой привязки отправителя запроса к его получателю, передавая
запрос по цепочке потенциальных обработчиков.
Применимость:
	- заранее неизвестен нужный обработчик;
	- есть необходимость задавать набор обработчиков динамически;
	- есть необходимость отправить запрос одному из нескольких обработчиков,
		явно этого не указывая.
Плюсы:
	- реализует принцип открытости/закрытости;
	- реализует принцип единой ответственности;
	- убирает жёсткую привязку отправителя к обработчику.
Минусы:
	- создаём дополнительные типы. иногда проще роутить через if-else/switch
*/

package main

import "fmt"

type handler interface {
	sendRequest(string) string
}

type handlerA struct {
	nextHandler handler
}

func (ha *handlerA) sendRequest(msg string) string {
	switch {
	case msg == "A":
		return fmt.Sprintf("Handler A done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

type handlerB struct {
	nextHandler handler
}

func (ha *handlerB) sendRequest(msg string) string {
	switch {
	case msg == "B":
		return fmt.Sprintf("Handler B done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

type handlerC struct {
	nextHandler handler
}

func (ha *handlerC) sendRequest(msg string) string {
	switch {
	case msg == "C":
		return fmt.Sprintf("Handler C done with %s", msg)
	case ha.nextHandler != nil:
		return ha.nextHandler.sendRequest(msg)
	default:
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

func main() {
	a := &handlerA{}
	b := &handlerB{}
	c := &handlerC{}
	a.nextHandler = b
	b.nextHandler = c

	req := "B"
	fmt.Println(a.sendRequest(req))

	req = "C"
	fmt.Println(a.sendRequest(req))

	req = "P"
	fmt.Println(a.sendRequest(req))
}
