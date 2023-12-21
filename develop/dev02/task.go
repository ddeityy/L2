package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	res, err := unpack(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func unpack(input string) (string, error) {
	runes := []rune(input)
	var b strings.Builder

	for i, v := range runes {
		if unicode.IsDigit(runes[i]) {
			if i == 0 {
				return "", fmt.Errorf("invalid string")
			}

			if unicode.IsDigit(runes[i-1]) {
				continue
			}

			letter := runes[i-1]

			num := strings.Builder{}
			num.WriteRune(v)

			for j := i + 1; j < len(runes); j++ {
				if !unicode.IsDigit(runes[j]) {
					break
				}
				num.WriteRune(runes[j])
			}

			repeats, err := strconv.Atoi(num.String())
			if err != nil {
				return "", err
			}

			for j := 0; j < repeats-1; j++ {
				b.WriteRune(letter)
			}

			continue
		}
		_, err := b.WriteRune(runes[i])
		if err != nil {
			return "", err
		}
	}

	return b.String(), nil
}
