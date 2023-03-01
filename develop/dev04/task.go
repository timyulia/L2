package main

import (
	"fmt"
	"sort"
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

func sortSym(word string) string {
	symbols := []byte(word)
	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i] < symbols[j]
	})
	return string(symbols)
}

func toLow(words []string) {
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}
}

func devide(words []string) map[string][]string {
	mapSorted := make(map[string]string) //ключ - слово с ортсорт символами, значение - первое такое слово
	result := make(map[string][]string)
	toLow(words)
	sort.Strings(words)

	for _, word := range words {
		sortedSyms := sortSym(word)
		if value, ok := mapSorted[sortedSyms]; ok {
			result[value] = append(result[value], word)
		} else {
			mapSorted[sortedSyms] = word
		}
	}
	for key := range result {
		result[key] = append(result[key], key)
	}
	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "том", "мот"}
	fmt.Println(devide(words))
}
