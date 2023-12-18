package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Команда — это поведенческий паттерн, позволяющий заворачивать
запросы или простые операции в отдельные объекты

Применение:
    Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети.
    Когда нужна операция отмены.

Плюсы:
    Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
    Позволяет реализовать простую отмену и повтор операций.
    Позволяет реализовать отложенный запуск операций.
    Позволяет собирать сложные команды из простых.

Минусы:
    Усложняет код программы из-за введения множества дополнительных классов.

Примеры:
	Многоуровневая отмена операций (undo)
	Индикаторы выполнения
*/

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Device interface {
	on()
	off()
}

type On struct {
	device Device
}

func (c *On) execute() {
	c.device.on()
}

type Off struct {
	device Device
}

func (c *Off) execute() {
	c.device.off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}
func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}
func main() {
	tv := &TV{}
	button := &Button{}

	on := &On{
		device: tv,
	}
	off := &Off{
		device: tv,
	}

	button.command = on
	button.press()

	button.command = off
	button.press()
}
