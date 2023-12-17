package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод — это порождающий паттерн проектирования,
который решает проблему создания различных продуктов,
без указания конкретных классов продуктов.

Применение:
	Отделяет код производства объектов от остального кода, который эти объекты использует.
    Позволяет экономить системные ресурсы путем повторного использования уже созданных объектов вместо порождения новых.

Плюсы:
    Избавляет главный класс от привязки к конкретным типам объектов.
    Выделяет код производства объектов в одно место, упрощая поддержку кода.
    Упрощает добавление новых типов объектов в программу.

Минусы:
    Может привести к созданию больших параллельных иерархий классов,
	так как для каждого типа объекта надо создать свой подкласс создателя.

Примеры:
	Создание подключений к разным базам данных в зависимости от запроса.
*/

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}
func (g *Gun) getName() string {
	return g.name
}
func (g *Gun) setPower(power int) {
	g.power = power
}
func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}

type handgun struct {
	Gun
}

func newHandgun() IGun {
	return &handgun{
		Gun: Gun{
			name:  "Handgun",
			power: 1,
		},
	}
}

func newGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "handgun" {
		return newHandgun(), nil
	}
	return nil, fmt.Errorf("Wrong gun type")
}

func main() {
	ak47, _ := newGun("ak47")
	handgun, _ := newGun("handgun")
	printDetails(ak47)
	printDetails(handgun)
}
func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
