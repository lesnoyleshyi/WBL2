/*
Паттерн фасад определяет интерфейс более высокого уровня вместо
набора интерфейсов некоторой подсистемы, тем самым упрощает ее использование.
Применимость:
- есть необходимость предоставить простой интерфейс к сложной подсистеме
или добавить независимости между клиентами и классами реализации
 Плюсы:
 - изолирует клиентов от сложной подсистемы
 - cохраняет возможность клиентам напрямую обращаться к классам подсистемы
 Минусы:
 - усложняет доступ к подсистеме если является единственной точкой доступа к ней
 - может стать монструозной конструкцией если привязать к нему слишком много функционала
*/

package main

import (
	"fmt"
)

type gpu struct {
}

func (gpu) StartGPU() {
	fmt.Println("GPU started")

}

type cpu struct {
}

func (cpu) StartCPU() {
	fmt.Println("CPU started")
}

type memory struct {
}

func (g *memory) InitMemory() {
	fmt.Println("Memory init")
}

type computerFacade struct {
	Memory memory
	GPU    gpu
	CPU    cpu
}

func (c computerFacade) startComputer() {
	c.Memory.InitMemory()
	c.CPU.StartCPU()
	c.GPU.StartGPU()
}

func main() {
	pc := computerFacade{}
	pc.startComputer()
}
