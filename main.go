package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	clues := generateClues(index, os.Args[1:])
	fmt.Println(clues)
}
