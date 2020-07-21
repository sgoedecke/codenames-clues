# Codenames Clues

Generate codenames clues from the command line.

## Installation

`go build`

## Usage

Make sure the `text/` folder has a bunch of source files in it. Then run `./codenames-clues [my clues]`, e.g. `./codenames-clues arthur table knight`. This will build an index and save it as `./index`, which will be used for subsequent runs. If you change the source text files, you will need to delete the index to have those changes picked up.

## TODO

Keep track of the frequency of intersections when calculating clues and prioritize high-intersection words
Filter out non-nouns in a nicer way than maintaining a common words list
