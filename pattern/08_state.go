package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние — это поведенческий паттерн, позволяющий динамически
изменять поведение объекта при смене его состояния.

Применение:
	Состояние нужно использовать в случаях, когда объект может иметь много
	различных состояний, которые он должен менять в зависимости
	от конкретного поступившего запроса.

Плюсы:
	Избавляет от множества больших условных операторов машины состояний.
	Концентрирует в одном месте код, связанный с определённым состоянием.

Минусы:
	Может неоправданно усложнить код, если состояний мало и они редко меняются

Примеры:
	Планировщик горутин выбирает что делать с горутиной в зависимости от её состояния.
*/

type Context struct {
	state State
}

func (c *Context) DoSmth() {
	c.state.Handle()
}

func (c *Context) SetState(state State) {
	c.state = state
}

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("State A")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("State B")
}

func main() {
	context := Context{new(ConcreteStateA)}
	context.DoSmth()
	context.SetState(new(ConcreteStateB))
	context.DoSmth()
}
