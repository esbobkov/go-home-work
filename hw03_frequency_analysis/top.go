package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

const defaultTopLength = 10

type wordCounter struct {
	word  string
	count int
}

func Top10(text string) []string {
	if text == "" {
		return make([]string, 0)
	}

	words := strings.Fields(text)
	wordsCounts := make(map[string]int)

	for _, v := range words {
		wordsCounts[v]++
	}

	res := make([]wordCounter, 0)

	for w, c := range wordsCounts {
		res = append(res, wordCounter{word: w, count: c})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].count > res[j].count
	})

	topLength := getTopLength(res)

	top := make([]string, topLength)
	for i := 0; i < topLength; i++ {
		top[i] = res[i].word
	}

	return top
}

func getTopLength(res []wordCounter) int {
	l := len(res)
	if defaultTopLength > l {
		return l
	}
	return defaultTopLength
}
