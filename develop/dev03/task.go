package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	sortColumn   int
	sortByNum    bool
	reversedSort bool
	uniqueValues bool
}

func parseFlags() *flags {
	s := flags{}

	flag.IntVar(&s.sortColumn, "k", -1, "column to sort")
	flagN := flag.Bool("n", false, "sort by number")
	flagR := flag.Bool("r", false, "reverse")
	flagU := flag.Bool("u", false, "return only unique values")

	flag.Parse()

	s.sortByNum = *flagN
	s.uniqueValues = *flagU
	s.reversedSort = *flagR

	return &s
}

func sorting(keys []string, numSort bool) []string {
	if numSort {
		nums := make([]string, 0)
		letters := make([]string, 0)
		for _, word := range keys {
			if _, err := strconv.Atoi(word); err == nil {
				nums = append(nums, word)
			} else {
				letters = append(letters, word)
			}
		}
		sort.Strings(letters)
		sort.Slice(nums, func(i, j int) bool {
			numA, _ := strconv.Atoi(nums[i])
			numB, _ := strconv.Atoi(nums[j])
			return numA < numB
		})
		keys = append(letters, nums...)
	} else {
		sort.Strings(keys)
	}
	return keys
}

func reverse(lines []string) {
	n := len(lines)
	for i := 0; i < n/2; i++ {
		lines[i], lines[n-i-1] = lines[n-i-1], lines[i]
	}
}

func makeUnique(lines []string) []string {
	unique := make(map[string]struct{})
	for _, line := range lines {
		unique[line] = struct{}{}
	}
	res := make([]string, 0)
	for key, _ := range unique {
		res = append(res, key)
	}
	return res
}

func sortLinesByColumnes(data []string, k int, numSort bool) []string {
	keys := make(map[string][]string)
	allKeys := make([]string, 0)
	for _, line := range data {
		words := strings.Split(line, " ")
		if len(words) < k {
			if _, ok := keys[""]; !ok {
				allKeys = append(allKeys, "")
			}
			keys[""] = append(keys[""], line)
		} else {
			if _, ok := keys[words[k-1]]; !ok {
				allKeys = append(allKeys, words[k-1])
			}
			keys[words[k-1]] = append(keys[words[k-1]], line)
		}
	}
	allKeys = sorting(allKeys, numSort)
	res := make([]string, 0)
	for _, key := range allKeys {
		for _, val := range keys[key] {
			res = append(res, val)
		}
	}
	return res
}

func sortLines(data []string, fl flags) []string {
	var res []string
	if fl.uniqueValues {
		data = makeUnique(data)
	}
	if fl.sortColumn != -1 {
		res = sortLinesByColumnes(data, fl.sortColumn, fl.sortByNum)
	} else {
		res = sorting(data, fl.sortByNum)
	}
	if fl.reversedSort {
		reverse(res)
	}
	return res
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

func writeToFile(filename string, data []string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(f)

	for i := 0; i < len(data); i++ {
		_, err := fmt.Fprintln(f, data[i])
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func makeSort(filename string) {
	flg := parseFlags()
	data := readFile(filename)
	dataSorted := sortLines(data, *flg)
	writeToFile("sorted.txt", dataSorted)
}

func main() {
	filename := "text.txt"
	makeSort(filename)
}
