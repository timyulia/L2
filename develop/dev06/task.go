package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	fields    map[int]struct{}
	delimiter string
	separated bool
}

// Функция парсит флаги из аргументов запуска программы
func parseFlags() *flags {
	f := flags{}

	var fieldList string
	flag.StringVar(&fieldList, "f", "", "choose fields")
	flag.StringVar(&f.delimiter, "d", "	", "delimiter")
	flagS := flag.Bool("s", false, "separated")

	flag.Parse()

	f.separated = *flagS

	var err error
	f.fields, err = parseFieldList(fieldList)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &f
}

func cut(fl flags) {
	output := make([][]string, 0)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		row, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad input")
		}
		if row == "quit\n" {
			for _, v := range output {
				fmt.Println(strings.Join(v, " "))
			}

			break
		}

		if fl.separated {
			if !strings.Contains(row, fl.delimiter) {
				continue
			}
		}
		row = strings.TrimSuffix(row, "\n")
		splitRow := strings.Split(row, fl.delimiter)

		if fl.fields != nil {
			filtredRow := make([]string, 0)
			for i, v := range splitRow {
				if _, ok := fl.fields[i+1]; ok {
					filtredRow = append(filtredRow, v)
				}
			}

			output = append(output, filtredRow)
		} else {
			output = append(output, splitRow)
		}
	}
}

func parseFieldList(fieldList string) (map[int]struct{}, error) {
	if fieldList == "" {
		return nil, nil
	}
	fieldList = strings.TrimSpace(fieldList)
	intStrList := strings.Split(fieldList, ",")
	res := make(map[int]struct{}, len(intStrList))
	for _, intStr := range intStrList {
		num, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, errors.New("bad field list")
		}
		res[num] = struct{}{}
	}
	return res, nil
}

func main() {
	f := flags{}
	f = *parseFlags()
	cut(f)
}
