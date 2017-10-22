package string

// Search looks for a pattern in the given text, returns true if founf, false otherwise
func Search(pattern, text string) bool {
	i, k := 0, len(pattern)
	table := make(map[string]int)

	for j, item := range pattern {
		table[string(item)] = k - j - 1
	}
	for a := 0; a < len(text); a++ {
		if text[i:i+k] == pattern {
			return true
		}
		if i+k < len(text) {
			if _, ok := table[string(text[i+k])]; ok {
				i += table[string(text[i+k])]
			} else {
				i = k + 1
			}
		}
	}

	return false
}
