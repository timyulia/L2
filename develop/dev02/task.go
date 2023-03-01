package main

import (
	"errors"
	"fmt"
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

var errDigitBegin = errors.New("incorrect string: begins with a digit")
var errLastEscape = errors.New("incorrect string: cannot have only one last escape")
var errAfterEscape = errors.New("incorrect string: there can be only a digit after escape")

func unpack(str string) (string, error) {
	str = strings.Trim(str, "\n")
	if str == "" {
		return str, nil
	}
	syms := []rune(str)
	prev := syms[0]
	if unicode.IsDigit(prev) {
		return "", errDigitBegin
	}
	var b strings.Builder
	b.WriteRune(prev)
	number := []rune{}
	fl := true
	i := 1
	for i < len(syms) {
		switch {
		case unicode.IsDigit(syms[i]) && fl:
			number = append(number, syms[i])
		case string(syms[i]) == `\`:
			if i == len(syms)-1 {
				return "", errLastEscape
			}
			if string(syms[i+1]) == `\` {
				prev = syms[i]
				b.WriteRune(prev)
				i++
			} else if unicode.IsDigit(syms[i+1]) {
				fl = false
			} else {
				return "", errAfterEscape
			}
		default:
			if len(number) > 0 {
				num, _ := strconv.Atoi(string(number))
				for j := 1; j < num; j++ {
					b.WriteRune(prev)
				}
			}
			prev = syms[i]
			b.WriteRune(prev)
			number = []rune{}
			fl = true
		}
		i++
	}
	if len(number) > 0 {
		num, _ := strconv.Atoi(string(number))
		for j := 1; j < num; j++ {
			b.WriteRune(prev)
		}
	}
	return b.String(), nil
}

func main() {
	var str string
	fmt.Println("enter exit to stop")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	for str != "exit" {
		res, err := unpack(str)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
		_, err = fmt.Scanln(&str)
		if err != nil {
			return
		}
	}
}
