package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func buildIndex(filePaths []string) map[string][]string {
	// build index
	index := make(map[string][]string)
	data := ""
	fmt.Println("Scanning files...")
	for _, path := range filePaths {
		dat, _ := ioutil.ReadFile(path)
		data = data + string(dat)
	}

	fmt.Println("Building index...")

	commonWords := []string{"by", "of", "for", "in", "and", "is", "from", "was", "the", "a", "an", "into", "as", "but", "with", "to", "who", "which", "out", "also", "each", "where", "than", "has", "that", "not", "on", "so", "no", "its", "would", "may", "began", "became", "gave", "till", "other", "his", "one", "two", "upon", "during", "it", "then", "after", "many", "de", "et", "came", "be", "there", "are", "all", "their", "went", "were", "some", "any", "very", "they", "when", "this", "again", "himself", "him", "her", "herself", "or", "defn", "over", "among", "had", "q", "both", "if", "our", "do", "have", "can", "your", "while", "now", "did", "we", "get", "now", "give", "do", "up", "such", "those", "etym", "akin", "obs", "vb", "these", "been", "w", "fr", "cf"}

	blocks := strings.Split(data, "\n")
	for _, block := range blocks {
		words := strings.Split(block, " ")

		for _, word := range words {
			// remove extra whitespace and treat semicolons like full stops
			re := regexp.MustCompile(`[^A-Za-z]+`)
			word = re.ReplaceAllString(word, "")
			word = strings.ToLower(word)

			if contains(commonWords, word) {
				continue
			}

			if index[word] == nil {
				index[word] = make([]string, 0)
			}

			for _, w := range words {
				w = re.ReplaceAllString(w, "")
				w = strings.ToLower(w)
				if contains(commonWords, w) {
					continue
				}
				index[word] = append(index[word], w)
			}
		}
	}

	fmt.Println("Index built!")
	return index
}

func generateClues(index map[string][]string, args []string) []string {
	clues := index[strings.ToLower(args[0])]

	for _, arg := range args {
		arg = strings.ToLower(arg)
		clues = intersections(clues, index[arg])
	}

	filteredClues := []string{}

	for _, clue := range clues {
		if !contains(args, clue) {
			filteredClues = append(filteredClues, clue)
		}
	}

	return filteredClues
}
