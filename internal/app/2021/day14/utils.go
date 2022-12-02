package day14

func SlidingWindowString(size int, input string) []string {
	// returns the input slice as the first element
	if len(input) <= size {
		return []string{input}
	}

	// allocate slice at the precise size we need
	r := make([]string, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
