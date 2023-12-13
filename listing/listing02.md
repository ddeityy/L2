Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
	2
	1
```

Правила работы `defer`:

1. Аргументы отложенной функции оцениваются при вызове оператора `defer`:
	```go
	func a() {
	    i := 0
	    defer fmt.Println(i) // выведет 0
	    i++
		defer fmt.Println(i) // выведет 1
	    return
	}
	a() // выведет 1 0 из-за второго правила
	```
2. Отложенные вызовы функций выполняются в порядке LIFO после возврата из основной функциию:
	```go
	func b() {
		for i := 0; i < 4; i++ {
			defer fmt.Print(i) // выведет 3210
		}
	}
	```
3. Отложенные функции имеют доступ к именованным возвращаемым значениям.
	```go
	func c() (i int) {
		defer func() { 
			fmt.Println(i) // выведет 1 т.к. `return 1` эквивалентно `i = 1; return i`
			i++
		}()
		return 1
	}
	```

Функция `test` попадает под третье правило, поэтому `defer` увеличит переменную. 

В функции `anotherTest` `defer` не имеет доступа к возвращаемой переменной `x`.
Соответственно в основной функции `x = 1` и это значение возвращается.