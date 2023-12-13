Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Интерфейс равен `nil`, только если и тип, и значение равны `nil`.
В нашем случае функция `test` возвращает интерфейс `error`, в котором данные будут `nil`,
но тип = `*main.customError`, и он не равен `nil`.

Это можно проверить с помощью библиотеки reflect.
```go
fmt.Println(reflect.TypeOf(err).String()) // *main.customError
fmt.Println(reflect.ValueOf(err).IsNil()) // true
```
