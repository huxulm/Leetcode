package uniquemorsecodewords

func uniqueMorseRepresentations(words []string) int {
	var t = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	var m = map[string]bool{}

	for _, w := range words {
		trans := ""
		for _, ch := range w {
			trans += t[ch-'a']
		}
		m[trans] = true
	}

	return len(m)
}
