package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Цепочка вызовов — это поведенческий паттерн, позволяющий
передавать запрос по цепочке потенциальных обработчиков, пока один
из них не обработает запрос.

Применение:
    Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
    Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
    Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы:
    Уменьшает зависимость между клиентом и обработчиками.
    Реализует принцип единственной обязанности.
    Реализует принцип открытости/закрытости.

Минусы:
    Запрос может остаться никем не обработанным.

Примеры:
	Обработка запроса веб сервера
*/

type IHandler interface {
	SetNext(IHandler) IHandler
	Handle(string) string
}

type Handler struct {
	next IHandler
}

func (h *Handler) SetNext(next IHandler) IHandler {
	h.next = next
	return next
}

func (h *Handler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

type FirstReceiver struct {
	Handler
}

func (f *FirstReceiver) Handle(request string) string {
	if request == "" {
		return "empty request handler"
	}
	return f.next.Handle(request)
}

func (f *FirstReceiver) setNext(next IHandler) IHandler {
	f.next = next
	return next
}

type SecondReceiver struct {
	Handler
}

func (s *SecondReceiver) Handle(request string) string {
	if request == "Hello" {
		return "hello handler"
	}
	return s.next.Handle(request)
}

func (s *SecondReceiver) setNext(next IHandler) IHandler {
	s.next = next
	return next
}

func main() {
	FirstReceiver := &FirstReceiver{}

	SecondReceiver := &SecondReceiver{}
	FirstReceiver.setNext(SecondReceiver)

	request := ""
	response := FirstReceiver.Handle(request)
	fmt.Println(response)

	request = "Hello"
	response = FirstReceiver.Handle(request)
	fmt.Println(response)

}
