package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadDictionary(path string) []string {
	bytes, _ := ioutil.ReadFile(path)
	return strings.Fields(strings.ToUpper(string(bytes)))
}

// Any information after a blank line is ignored. This lets you store notes and
// different versions below.
func loadGrid(path string) [][]byte {
	bytes, _ := ioutil.ReadFile(path)
	allLines := strings.Split(strings.ToUpper(string(bytes)), "\n")
	lines := []string{}
	for _, l := range allLines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			break
		}
		lines = append(lines, l)
	}
	xdim := len(lines[0])
	ydim := len(lines)
	grid := make([][]byte, xdim)
	for x := 0; x < xdim; x++ {
		grid[x] = make([]byte, ydim)
		for y := 0; y < ydim; y++ {
			grid[x][y] = lines[y][x]
		}
	}
	return grid
}

func printWord(word word, grid [][]byte, wordIndex int) {
	var str string
	for _, l := range word.letters {
		str += string(grid[l.x][l.y])
	}
	var ori string
	if word.orientation == down {
		ori = "down"
	} else {
		ori = "across"
	}
	if len(word.candidates) > 10 {
		fmt.Printf("%d. %s %5d %v  %6s\n",
			wordIndex,
			str,
			len(word.candidates),
			word.letters[0],
			ori)
	} else {
		fmt.Printf("%d. %s %5d %v  %6s %v\n",
			wordIndex,
			str,
			len(word.candidates),
			word.letters[0],
			ori,
			word.candidates)
	}
}

func printWords(words []word, grid [][]byte) {
	for i, w := range words {
		printWord(w, grid, i)
	}
}

func printGrid(grid [][]byte) {
	xdim := len(grid)
	ydim := len(grid[0])
	for y := 0; y < ydim; y++ {
		for x := 0; x < xdim; x++ {
			fmt.Printf("%c", grid[x][y])
		}
		fmt.Printf("\n")
	}
}

func prettyCandidates(words []word) string {
	cs := []string{}
	for _, w := range words {
		cs = append(cs, w.candidates[0])
	}
	return strings.Join(cs, " ")
}
