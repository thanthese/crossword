package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	across = iota
	down
)

// Pair represents grid coordinates. Upper left is (0, 0).
type coord struct {
	x int
	y int
}

// Word represents an answer on the puzzle -- 1 Down or 3 Across -- as well as
// lots of meta data about it necessary for calculating.
type word struct {
	orientation int

	// We can't just store the letters of the word because they're shared with
	// other words. Instead we store a "pointer" -- the grid coordinates -- to
	// each letter on the official grid.
	letters []coord

	// Each letter has a matching buddy.
	buddies []buddyPtr

	// List of words that could still possibly go in this spot.
	candidates []string
}

// Each square on the grid has two words intersecting -- "buddies". This struct
// is used as a pointer from one letter to its buddy letter in another word.
type buddyPtr struct {
	wordId   int // which word I go with, according to the main word[] slice
	letterId int // how far along in that word I go
}

type solution struct {
	words []word
	grid  [][]byte
}

func getOptions() (wordsFile, gridFile string, abbreviated bool) {
	words := flag.String("dict", "", "Words to use.")
	grid := flag.String("grid", "grid.txt", "Grid template to use.")
	abbr := flag.Bool("abbr", false, "Show abbreviated output.")
	flag.Parse()
	if _, err := os.Stat(*words); err != nil {
		fmt.Printf("ERROR: Cannot read --dict file \"%s\".\n", *words)
		flag.PrintDefaults()
		os.Exit(1)
	}
	if _, err := os.Stat(*grid); err != nil {
		fmt.Printf("ERROR: Cannot read --grid file \"%s\".\n", *grid)
		flag.PrintDefaults()
		os.Exit(1)
	}
	return *words, *grid, *abbr
}

func main() {
	var wordsFile, gridFile, abbr = getOptions()

	grid := loadGrid(gridFile)
	dictionary := loadDictionary(wordsFile)
	words := initWords(grid, dictionary)
	reduceCandidates(words, grid)

	solutions := make(chan solution)
	go func() {
		search(words, grid, solutions)
		close(solutions)
	}()

	var solutionsFound int = 0
	for s := range solutions {
		solutionsFound++
		fmt.Printf("%6d:  %s\n", solutionsFound, prettyCandidates(s.words))
		if !abbr {
			printGrid(s.grid)
			fmt.Println()
		}
	}
	fmt.Println("Solutions found:", solutionsFound)
}
