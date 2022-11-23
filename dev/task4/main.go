package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Anagram(str1, str2 string) bool {
	if utf8.RuneCountInString(str1) != utf8.RuneCountInString(str2) {
		return false
	}
	m := make(map[string]int)

	for _, v := range strings.Trim(strings.ToLower(str1), " ") {
		m[string(v)]++
	}

	for _, v := range strings.Trim(strings.ToLower(str2), " ") {
		m[string(v)] = m[string(v)] - 1
		if m[string(v)] < 0 {
			return false
		}
	}

	return true
}

func SaveAnagram(anagrams []string) map[string][]string {
	m := make(map[string][]string)

	for _, anagram := range anagrams {
		s := []string{}

		for _, word := range anagrams {
			if Anagram(anagram, word) {
				s = append(s, strings.ToLower(word))
				if len(s) > 1 {
					m[s[0]] = s[1:]
				}
			}
		}
	}

	for k := range m {
		sort.Strings(m[k])
	}

	return m
}

func main() {
	mp := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}
	fmt.Println(SaveAnagram(mp))
}
