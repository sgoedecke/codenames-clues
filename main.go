package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	var index map[string][]string
	b := new(bytes.Buffer)

	if fileExists("index") {
		// read index from file
		d := gob.NewDecoder(b)
		d.Decode(&index)
	} else {
		// build and serialize index to file
		sourceFiles, _ := filepath.Glob("./text/*")
		index = buildIndex(sourceFiles)

		e := gob.NewEncoder(b)
		e.Encode(index)

		ioutil.WriteFile("index", b.Bytes(), 0644)
	}

	clues := generateClues(index, os.Args[1:])
	fmt.Println(clues)
}
