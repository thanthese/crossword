package main

// For each word, throw out candidates that no longer fit given the new
// realities of the grid.
func reduceCandidates(words []word, grid [][]byte) {
	for keepChecking := true; keepChecking; {
		keepChecking = false
		for i, w := range words {
			candidates := []string{}
			for _, c := range w.candidates {
				if candidateFits(c, w, words, grid) {
					candidates = append(candidates, c)
				} else {
					keepChecking = true
				}
			}
			words[i].candidates = candidates
		}
	}
}

// There are two types of fitting, and a candidate must pass both.
//
// 1. Does the candidate fit into the current state of the grid? Meaning, if
// the grid already has a letter there it must be the same letter as the
// candidate has at that spot.
//
// 2. Does the candidate have a match amongst the options for the words going
// the other way (among the buddies)?
func candidateFits(candidate string, word word, words []word, grid [][]byte) bool {
	for i, coord := range word.letters {
		letter := grid[coord.x][coord.y]
		if candidate[i] != letter && letter != '.' {
			return false
		}
		buddyPtr := word.buddies[i]
		buddyWords := words[buddyPtr.wordId].candidates
		if !letterMatches(candidate[i], buddyWords, buddyPtr.letterId) {
			return false
		}
	}
	return true
}

// Does the letter ever match the words at the given index?
func letterMatches(letter byte, words []string, slot int) bool {
	for _, w := range words {
		if letter == w[slot] {
			return true
		}
	}
	return false
}
