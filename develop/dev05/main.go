package main

import (
	"fmt"
	"grep/grep"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

func main() {
	var app grep.App
	err := app.FromArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}
	res, err := app.Grep()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
