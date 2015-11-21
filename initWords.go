package main

func initWords(grid [][]byte, dict []string) []word {
	words := buildCoords(grid)
	addBuddies(words)
	addCandidates(words, dict)
	return words
}

// Create the words and add the fundamental, coordinate-related pieces: letters
// and orientation.
func buildCoords(grid [][]byte) []word {
	xdim := len(grid)
	ydim := len(grid[0])
	words := []word{}
	for x := 0; x < xdim; x++ {
		for y := 0; y < ydim; y++ {

			if grid[x][y] == '#' {
				continue
			}

			// beginning of across word
			if x == 0 || grid[x-1][y] == '#' {
				w := word{}
				w.orientation = across
				for ix := x; ix < xdim; ix++ {
					if grid[ix][y] == '#' {
						break
					}
					w.letters = append(w.letters, coord{ix, y})
				}
				words = append(words, w)
			}

			// beginning of down word
			if y == 0 || grid[x][y-1] == '#' {
				w := word{}
				w.orientation = down
				for iy := y; iy < ydim; iy++ {
					if grid[x][iy] == '#' {
						break
					}
					w.letters = append(w.letters, coord{x, iy})
				}
				words = append(words, w)
			}
		}
	}
	return words
}

func addBuddies(words []word) {
	for i, w := range words {
		words[i].buddies = make([]buddyPtr, len(w.letters))
		for j, let := range w.letters {
			for k, iw := range words {
				if w.orientation == iw.orientation {
					continue
				}
				for l, ilet := range iw.letters {
					if let == ilet {
						words[i].buddies[j].wordId = k
						words[i].buddies[j].letterId = l
					}
				}
			}
		}
	}
}

// Add candidates to each word's possible word list based on only the most
// rudimentary filtering -- by length. This is a small optimization so we're
// not lugging around several copies of the full dictionary.
func addCandidates(words []word, dictionary []string) {
	for i := range words {
		candidates := []string{}
		for _, d := range dictionary {
			if len(words[i].letters) == len(d) {
				candidates = append(candidates, d)
			}
		}
		words[i].candidates = candidates
	}
}
