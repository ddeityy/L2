Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
```
В программе возникает дедлок из-за того, что мы пытаемся прочитать из канала, в который перестали писать, но не закрыли.
Чтобы исправить эту ситуацию, нужно закрыть канал в конце горутине.
```go
go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // <-
	}()
```