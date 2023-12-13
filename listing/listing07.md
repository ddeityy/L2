Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
0
0
...
```

В функции `merge` в блоке `select` нет проверки на то,
возвращается ли из канала стандартное значение.
После закрытия входного канала (`a` и `b`) в выходной канал будут отправляться нули (стандартное значение для `int`).

В Go нет способа проверить закрыт ли канал не вычитывая из него значения,
поэтому для слияния двух каналов нужно применить другой подход:


```go
func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func(ch <-chan int) {
		for v := range ch {
			out <- v
		}
		wg.Done()
	}(a)

	go func(ch <-chan int) {
		for v := range ch {
			out <- v
		}
		wg.Done()
	}(b)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```