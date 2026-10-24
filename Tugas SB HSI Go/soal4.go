package main

import (
	"fmt"
	"strings"
)

func calculateFreqs(word string) map[string]int {
	frekuensi := make(map[string]int)

	kataKata := strings.Split(word, " ")

	for _, kata := range kataKata {
		frekuensi[kata]++
	}
	return frekuensi

}

func main() {
	fmt.Println(calculateFreqs("i love vibe coding"))
	fmt.Println(calculateFreqs(""))
	fmt.Println(calculateFreqs("go go go"))
	fmt.Println(calculateFreqs("red blue red"))
	fmt.Println(calculateFreqs("one two one two"))
	fmt.Println(calculateFreqs("hello world hello go"))

}
