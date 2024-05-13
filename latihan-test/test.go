package main

import (
	"regexp"
	"sort"
	"strings"
)

func removeChar(sentence string, char byte) string {
	var result string
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != char {
			result += string(sentence[i])
		}
	}
	return result
}

func isPalindrome(sentence string) bool {
	if len(sentence) <= 0 {
		return true
	}
	// remove all symbol and space on string
	pattern := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	cleanedSentence := pattern.ReplaceAllString(strings.ReplaceAll(strings.ToLower(sentence), " ", ""), "")

	startIndex, endIndex := 0, len(cleanedSentence)-1
	for startIndex < endIndex {
		if cleanedSentence[startIndex] != cleanedSentence[endIndex] {
			return false
		}
		startIndex++
		endIndex--
	}
	return true
}

func findFirstNonRepeatedChar(sentence string) string {
	counter := make(map[rune]int)
	for _, v := range sentence {
		counter[v]++
	}
	for _, v := range sentence {
		if counter[v] == 1 {
			return string(v)
		}
	}
	return ""
}

func secondHighest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	sort.Ints(arr)
	return arr[len(arr)-2]
}
