package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	rExp       string
	filename   string
}

func parseFlags() *flags {
	f := flags{}

	flag.IntVar(&f.after, "A", 0, "after")
	flag.IntVar(&f.before, "B", 0, "before")
	flag.IntVar(&f.context, "C", 0, "after and before")
	flagC := flag.Bool("c", false, "number of lines")
	flagI := flag.Bool("i", false, "ignore case")
	flagV := flag.Bool("v", false, "invert")
	flagF := flag.Bool("F", false, "fixed")
	flagN := flag.Bool("n", false, "line number")

	flag.Parse()

	f.count = *flagC
	f.invert = *flagV
	f.ignoreCase = *flagI
	f.fixed = *flagF
	f.lineNum = *flagN

	return &f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func grep(pattern string, input []string, fl flags) {
	if fl.ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	res := make(map[int]struct{})
	for i, v := range input {
		str := v
		if fl.ignoreCase {
			str = strings.ToLower(str)
		}

		var found bool
		if fl.fixed {
			found = str == pattern
		} else {
			found = strings.Contains(str, pattern)
		}
		if !found == fl.invert {
			res[i] = struct{}{}
		}
	}
	if fl.count {
		fmt.Println(len(res))
		return
	}

	before := max(fl.before, fl.context)
	after := max(fl.after, fl.context)
	around(res, len(input), before, after)
	output(res, input, fl)
}

func output(res map[int]struct{}, input []string, fl flags) {
	sorted := make([]int, 0)
	for key := range res {
		sorted = append(sorted, key)
	}
	sort.Ints(sorted)
	for _, key := range sorted {
		var v strings.Builder
		if fl.lineNum {
			v.WriteString(fmt.Sprintf("%d:", key))
		}
		v.WriteString(input[key])
		fmt.Println(v.String())
	}
}

func around(res map[int]struct{}, n, before, after int) {
	for key := range res {
		for i := 1; i <= before; i++ {
			if key-i >= 0 {
				res[key-i] = struct{}{}
			}
		}
		for i := 1; i <= after; i++ {
			if key+i < n {
				res[key+i] = struct{}{}
			}
		}
	}
}

func readFile(filename string) []string {
	var rows []string

	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(file)

	if err != nil {
		log.Fatal(err.Error(), filename)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}

	return rows
}

func main() {
	flg := parseFlags()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: [flags] [pattern] [file]")
		os.Exit(1)
	}
	pattern := args[0]
	filename := args[1]
	data := readFile(filename)
	grep(pattern, data, *flg)
}
