package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortRegister(s []string) []string {
	sort.SliceStable(s, func(i, j int) bool {
		return strings.ToLower(sl[i]) < strings.ToLower(sl[j])
	})

	return s
}

func index(s string, w []string) int {
	for i, v := range w {
		if s == v {
			return i
		}
	}

	return -1
}

func readScan(scan *bufio.Scanner) []string {
	s := []string{}

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

func sortDouble(sl []string) []string {

	s := make([]string, 0)

	for _, v := range sl {
		if index(v, s) < 0 {
			s = append(s, v)
		}
	}

	// возврат слайса отсортированного
	return sortRegister(s)
}

func sortReverse(sl []string) []string {

	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	// возврат слайса отсортированного
	return sl
}

func sortColumn(lines []string, k int, n bool) []string {

	s := make([][]string, 0)

	k = k - 1
	if k < 0 {
		k = 0
	}

	for _, l := range lines {
		s = append(s, strings.Split(l, " "))
	}

	if n {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				x, err := strconv.Atoi(s[i][k])
				y, err := strconv.Atoi(s[j][k])
				if err != nil {
					fmt.Println(err)
					return false
				}

				return x < y
			}

			return false
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				return strings.ToLower(s[i][k]) < strings.ToLower(s[j][k])
			}
			return false
		})
	}

	var str string
	sl = make([]string, 0)
	for _, l := range s {
		str = strings.Join(l, " ")
		sl = append(sl, str)
	}

	// возвращаем уже отсортированный слайс
	return sl
}

func unixSort(sl []string, flags *FlagsSort) []byte {
	sl = sortRegister(sl)

	// сортировка с удалением дублей
	if flags.unique {
		sl = sortDouble(sl)
	}

	// сортировка по колонке
	if flags.column > -1 {
		sl = sortColumn(sl, flags.column, flags.byName)
	}

	// сортировка в обратном порядке
	if flags.reverse {
		sl = sortReverse(sl)
	}

	return []byte(strings.Join(sl, "\n"))
}

const (
	DefaultColumnVal = -1
)

var fscan *bufio.Scanner
var fileName string
var column int
var Num bool
var unique bool
var reverse bool
var sl []string

type FlagsSort struct {
	column  int
	reverse bool
	unique  bool
	byName  bool
}

func main() {
	flag.IntVar(&column, "k", -1, "указание колонки для сортировки")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&Num, "n", false, "сортировать по числовому значению")
	flag.Parse()

	fileName = flag.Arg(0)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fl := &FlagsSort{unique: unique, column: column, reverse: reverse, byName: Num}
	fscan = bufio.NewScanner(f)
	sl = readScan(fscan)

	err = ioutil.WriteFile(f.Name(), unixSort(sl, fl), fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
