Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
```

Несмотря на то, что базовая структура равна `nil`, сам интерфейс не равен `nil`.

Интерфейс равен `nil`, только если и тип, и значение равны `nil`.

## Устройство интерфейсов
**Структура `iface`**

`iface` — это корневой тип, представляющий интерфейс рантайма ([src/runtime/runtime2.go](https://github.com/golang/go/blob/e822b1e26e20ef1c76672c0b77b0fd8a97a1fe84/src/runtime/runtime2.go#L202)).
Его определение выглядит так:

```go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}
```

Интерфейс представляет собой структуру, которая содержит 2 указателя:
- `tab` содержит адрес объекта `itable` — структуру, которая хранит некоторые метаданные о типе и список методов, используемых для удовлетворения интерфейса.
- `data` указывает на фактическую переменную с конкретным (статическим) типом.

Поскольку интерфейсы могут содержать только указатели, любое конкретное значение, которое мы переносим в интерфейс, должно иметь адрес.

**Структура `itab`**

`itab` определяется следующим образом ([src/runtime/runtime2.go](https://github.com/golang/go/blob/e822b1e26e20ef1c76672c0b77b0fd8a97a1fe84/src/runtime/runtime2.go#L902=)):

```go
type itab struct {
    inter *interfacetype
    _type *_type
    hash  uint32
    _     [4]byte
    fun   [1]uintptr
}
```

- `interfacetype` — это просто оболочка вокруг `_type` с некоторыми дополнительными метаданными, специфичными для интерфейса.

В текущей реализации эти метаданные в основном состоят из списка смещений, указывающих на соответствующие имена и типы методов, предоставляемых интерфейсом (`[]imethod`).

- `_type` описывает все аспекты типа: его имя, его характеристики (размер, выравнивание) и его поведение (сравнение, хеширование).

- `fun` содержит указатели на функции, составляющие виртуальную таблицу интерфейса.

Компилятор генерирует метаданные для каждого статического типа, в которых, помимо прочего, хранится список методов, реализованных для данного типа.

Рантайм вычисляет `Itab`, ища каждый метод, указанный в таблице методов типа интерфейса, в таблице методов конкретного типа. Рантайм кэширует itable после его создания, так что это соответствие нужно вычислить только один раз.

**Пустой интерфейс**

Структура данных для пустого интерфейса — это `iface` без `itab`. Тому есть две причины:

- Поскольку в пустом интерфейсе нет методов, все, что связано с динамической диспетчеризацией, можно смело выкинуть из структуры данных.
- Когда виртуальная таблица исчезла, тип самого пустого интерфейса всегда один и тот же.

`eface` — это корневой тип, представляющий пустой интерфейс в рантайме ([src/runtime/runtime2.go](https://github.com/golang/go/blob/e822b1e26e20ef1c76672c0b77b0fd8a97a1fe84/src/runtime/runtime2.go#L207)).
Его определение выглядит так:

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

Где `_type` содержит информацию о типе значения, на которое указывают данные.

Отличие пустого интерфейса от обычного заключается в отсутствии поля `itab`.
Поскольку у пустого интерфейса нет никаких методов, то и `itab` для него просчитывать и хранить не нужно — достаточно только метаинформации о статическом типе.