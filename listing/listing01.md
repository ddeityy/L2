Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
[77 78 79]
```
Создается массив `a` из 5 элементов, затем из этого массива делается слайс `b[1:4]` с `cap(b) = 4` и `len(b) = 3`.
В результате в слайсе будут элементы `[77 78 79]`.
Слайс `b` ссылается на массив `a`, и если изменить что-то в массиве `a`, то эти изменения будут видны в слайсе `b`.
Например:

```go
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4]
a[2] = 99
fmt.Println(b[1]) // 99
```

При создании слайса можно контролировать его ёмкость.
Ёмкость слайса должна быть <= ёмкости массива.
Например:

```go
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4:5]  //cap = 5
var c []int = a[1:4:6]  //ошибка
```