package main

import (
	"fmt"
	"reflect"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagrams(words []string) (map[string][]string, error) {
	if words == nil {
		return nil, fmt.Errorf("input slice must not be nil")
	}

	if len(words) < 2 {
		return nil, fmt.Errorf("input slice must have at least 2 elements")
	}

	result := make(map[string][]string)
	exists := make(map[string]struct{})

	for i, word := range words {
		exists[word] = struct{}{}
		for j := i + 1; j < len(words); j++ {
			if _, ok := exists[words[j]]; ok {
				continue
			}

			if anagrams(word, words[j]) {
				result[word] = append(result[word], words[j])
				exists[words[j]] = struct{}{}
			}
		}
	}

	return result, nil
}

// Приводим все строки к нижнему регистру
func formatData(words []string) []string {
	res := make([]string, 0, len(words))

	for _, v := range words {
		str := strings.ToLower(v)
		res = append(res, str)
	}

	return res
}

// Проверяем если 2 строки - анаграмы
func anagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	s1 := make(map[rune]struct{})
	for _, v := range str1 {
		s1[v] = struct{}{}
	}

	s2 := make(map[rune]struct{})
	for _, v := range str2 {
		s2[v] = struct{}{}
	}

	return reflect.DeepEqual(s1, s2)
}

func main() {
	var words []string
	res, err := findAnagrams(words)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
