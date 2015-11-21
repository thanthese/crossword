package main

func copyWords(words []word) []word {
	out := make([]word, len(words))
	for i, w := range words {
		out[i].orientation = w.orientation
		out[i].letters = w.letters
		out[i].buddies = w.buddies
		out[i].candidates = append([]string{}, w.candidates...)
	}
	return out
}

func copyGrid(grid [][]byte) [][]byte {
	out := make([][]byte, len(grid))
	for x := 0; x < len(grid); x++ {
		out[x] = append([]byte{}, grid[x]...)
	}
	return out
}

// Return the index of the word with the least number of remaining viable
// candidate options (meaning at least 2), or -1 if no such word found. I
// believe that this is a key optimization.
func leastActiveCandidates(words []word) int {
	index := -1
	for i, w := range words {
		if len(w.candidates) <= 1 {
			continue
		}
		if index == -1 {
			index = i
			continue
		}
		if len(w.candidates) < len(words[index].candidates) {
			index = i
		}
	}
	return index
}

func foundDeadEnd(words []word) bool {
	return len(words[0].candidates) == 0
}

func foundSolution(words []word) bool {
	for _, w := range words {
		if len(w.candidates) != 1 {
			return false
		}
	}
	return true
}

// Main recursive search algorithm.
func search(words []word, grid [][]byte, solutions chan solution) {
	if foundDeadEnd(words) {
		return
	}
	if foundSolution(words) {
		completeSolution(words, grid)
		solutions <- solution{words, grid}
		return
	}
	index := leastActiveCandidates(words)
	for _, candidate := range words[index].candidates {
		w := copyWords(words)
		g := copyGrid(grid)
		setWord(candidate, w[index], g)
		reduceCandidates(w, g)
		search(w, g, solutions)
	}
}

// Fill in the grid with str at position word.
func setWord(str string, word word, grid [][]byte) {
	for i, coord := range word.letters {
		grid[coord.x][coord.y] = str[i]
	}
}

// A unique solution has been found, but the grid hasn't been completely filled
// out. Do that now.
func completeSolution(words []word, grid [][]byte) {
	// aggressively not worried about efficiency
	for _, w := range words {
		for i := range w.candidates[0] {
			grid[w.letters[i].x][w.letters[i].y] = w.candidates[0][i]
		}
	}
}
