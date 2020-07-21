package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"path/filepath"
    "flag"
    "strings"
)

func main() {
	var index map[string][]string

	if fileExists("index") {
		fmt.Println("Loading index from file...")
		data, _ := ioutil.ReadFile("index")
		d := gob.NewDecoder(bytes.NewBuffer(data))
		d.Decode(&index)
		fmt.Println("Index loaded!")
	} else {
		// build and serialize index to file
		sourceFiles, _ := filepath.Glob("./text/*")
		index = buildIndex(sourceFiles)

		b := new(bytes.Buffer)
		e := gob.NewEncoder(b)
		e.Encode(index)

		ioutil.WriteFile("index", b.Bytes(), 0644)
	}

    wordsPtr := flag.String("w", "", "words to guess, comma separated")
    avoidPtr := flag.String("a", "", "words to avoid, comma separated")
    flag.Parse()

	clues := solve(index, strings.Split(*wordsPtr, ","), strings.Split(*avoidPtr, ","))
	fmt.Println(clues)
}
