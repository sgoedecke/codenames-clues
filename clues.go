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
	for _, path := range filePaths {
		fmt.Println("Scanning file: " + path)
		dat, _ := ioutil.ReadFile(path)
		data = data + string(dat)
	}

	fmt.Println("Building index...")

	blocks := strings.Split(data, "\n")
	for _, block := range blocks {
		words := strings.Split(block, " ")

		for _, word := range words {
			// remove extra whitespace and treat semicolons like full stops
			re := regexp.MustCompile(`[^A-Za-z]+`)
			word = re.ReplaceAllString(word, "")
			word = strings.ToLower(word)

			if index[word] == nil {
				index[word] = make([]string, 0)
			}

			for _, w := range words {
				w = re.ReplaceAllString(w, "")
				w = strings.ToLower(w)
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

	commonWords := []string{"by", "of", "for", "in", "and", "is", "from", "was", "the", "a", "an", "into", "as", "but", "with", "to", "who", "which", "out", "also", "each", "where", "than", "has", "that", "not", "on", "so", "no", "its", "would", "may", "began", "became", "gave", "till", "other", "his", "one", "two", "upon", "during", "it", "then", "after", "many", "de", "et", "came", "be"}

	filteredClues := []string{}
	for _, clue := range clues {
		if !contains(commonWords, clue) {
			filteredClues = append(filteredClues, clue)
		}
	}

	return filteredClues
}
