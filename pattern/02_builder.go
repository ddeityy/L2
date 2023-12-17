package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель — это порождающий паттерн проектирования, который
позволяет создавать объекты пошагово.

Применимость:
    Если создание нескольких представлений объекта состоит из одинаковых этапов, которые отличаются в деталях.
    Если Вам нужно собирать сложные составные объекты, например, деревья Компоновщика.

Плюсы:
    Позволяет создавать объекты пошагово.
    Позволяет использовать один и тот же код для создания различных объектов.
    Изолирует сложный код сборки объекта от его основной бизнес-логики.

Минусы:
    Усложняет код программы из-за введения дополнительных классов.
    Клиент может оказаться привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть метода получения результата.

Прмиеры:
	Создание логгеров разного уровня
*/

type PC struct {
	RAM int
	CPU string
	GPU string
}

type PCBuilderI interface {
	setRAM(val int) PCBuilderI
	setCPU(val string) PCBuilderI
	setGPU(val string) PCBuilderI
	Build() *PC
}

type PCBuilder struct {
	ram int
	cpu string
	gpu string
}

func (b *PCBuilder) setRAM(val int) PCBuilderI {
	b.ram = val
	return b
}
func (b *PCBuilder) setCPU(val string) PCBuilderI {
	b.cpu = val
	return b
}
func (b *PCBuilder) setGPU(val string) PCBuilderI {
	b.gpu = val
	return b
}
func (b *PCBuilder) Build() *PC {
	if b.ram == 0 {
		b.ram = 2
	}
	if b.cpu == "" {
		b.cpu = "Default CPU"
	}
	if b.gpu == "" {
		b.gpu = "Default GPU"
	}
	return &PC{b.ram, b.cpu, b.gpu}
}

func newPCBuilder() PCBuilderI {
	return &PCBuilder{}
}

func main() {
	customPCBuilder := newPCBuilder()
	customPCBuilder.setCPU("Cusom CPU").setGPU("Custom GPU").setRAM(16).Build()
	fmt.Println(customPCBuilder)
	defaultPCBuilder2 := newPCBuilder().Build()
	fmt.Println(defaultPCBuilder2)
}
