package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	dict := []string{"Пятак", "пятка", "тяПка", "Листок", "слиток", "Столик", "123", "132"}
	fmt.Println(anagrams(dict)) //nolint:forbidigo
}

func letters(word string) map[rune]struct{} {
	letters := make(map[rune]struct{})
	for _, l := range word {
		letters[l] = struct{}{}
	}

	return letters
}

func contains(words []string, word string) bool {
	for _, a := range words {
		if a == word {
			return true
		}
	}
	return false
}

func anagrams(words []string) *map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		hasSet := false
		word = strings.ToLower(word)

		for key := range anagrams {
			if reflect.DeepEqual(letters(word), letters(anagrams[key][0])) {
				hasSet = true
				if !contains(anagrams[key], word) {
					anagrams[key] = append(anagrams[key], word)
				}
			}
			sort.Strings(anagrams[key])
		}
		if !hasSet {
			anagrams[word] = append(anagrams[word], word)
		}
	}
	return &anagrams
}
