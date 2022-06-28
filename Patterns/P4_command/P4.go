/*
Паттерн команда позволяет превращать запросы в объекты и
передавать их как аргументы при вызове методов, ставить в очередь, логировать,
 а также отменять.
Применимость:
 - когда запрос в виде объекта дает преимущества, например,
  возможность отменять или ставить в очередь.
Плюсы :
 - убирает прямую зависимость между объектами порождающими и выполняющими запросы
  - логирование, отмена, повтор, очередь, конструирование сложных запросов
Минусы:
 - усложняет код в связи с введением дополнительных классов под разные типы запросов
*/

package main

import "fmt"

type command interface {
	execute() string
}

type turnOnCommand struct {
	receiver *receiver
}

func (t *turnOnCommand) execute() string {
	return t.receiver.turnOn()
}

type turnOffCommand struct {
	receiver *receiver
}

func (t *turnOffCommand) execute() string {
	return t.receiver.turnOff()
}

type receiver struct {
}

func (r *receiver) turnOn() string {
	return "Turning ON"
}

func (r *receiver) turnOff() string {
	return "Turning OFF"
}

type invoker struct {
	commands []command
}

func (i *invoker) storeCommand(cmd command) {
	i.commands = append(i.commands, cmd)
}

func (i *invoker) resetCommands(cmd command) {
	i.commands = make([]command, 0)
}

func (i *invoker) ExecCommand() string {
	var res string
	for _, cmd := range i.commands {
		res += cmd.execute() + "\n"
	}
	return res
}

func main() {
	inv := &invoker{}

	cmdON := &turnOnCommand{}
	cmdOFF := &turnOffCommand{}

	inv.commands = append(inv.commands, cmdON, cmdOFF, cmdON)
	fmt.Print(inv.ExecCommand())
}
