/*
Паттерн состояние позволяет объекту варьировать свое поведение в зависимости от внутреннего
состояния. Извне создается впечатление, что изменился класс объекта.
Применимость:
 - когда поведение объекта зависит от его состояния и должно изменяться во
   время выполнения
 - когда в коде операций встречаются состоящие из многих ветвей условные
   операторы, в которых выбор ветви зависит от состояния.
Плюсы:
 - избавляет от множества больших условных операторов машины состояний
 - концентрирует в одном месте код, связанный с определённым состоянием
 -  упрощает код контекста.
 Минусы:
 - может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

package main

import "fmt"

type exchangeStater interface {
	balance() string
}

type moscowExchange struct {
	balance int
	state   exchangeStater
}

func new(balance int, state exchangeStater) moscowExchange {
	return moscowExchange{balance: balance, state: state}
}

type openedExchange struct {
	exchange moscowExchange
}

func (op *openedExchange) balance() string {
	return fmt.Sprintf("Exchange is opened. Current balance is %d", op.exchange.balance)
}

type closedExchange struct {
	exchange moscowExchange
}

func (cl *closedExchange) balance() string {
	return fmt.Sprintf("Exchange is closed. Balance is fixed at %d", cl.exchange.balance)
}

func main() {
	ex := new(228, nil)
	op := &openedExchange{exchange: ex}
	cl := &closedExchange{exchange: ex}

	ex.state = op

	fmt.Println(ex.state.balance())

	ex.state = cl
	fmt.Println(ex.state.balance())
}
