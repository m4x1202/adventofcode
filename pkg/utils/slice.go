package utils

func SlidingWindow[S ~[]E, E ~[]X, X any](slice E, windowSize int) (window S) {
	// returns the input slice as the first element
	if len(slice) <= windowSize {
		return S{slice}
	}

	// allocate slice at the precise size we need
	r := make(S, 0, len(slice)-windowSize+1)

	for i, j := 0, windowSize; j <= len(slice); i, j = i+1, j+1 {
		r = append(r, slice[i:j])
	}

	return r
}

func CombineFunc[S ~[]E, E ~[]X, X any](s S, comb func(E) X) E {
	res := make(E, len(s))
	for i, v := range s {
		res[i] = comb(v)
	}
	return res
}

func SliceMap[S ~[]E, E any](s S, f func(E) E) S {
	if len(s) < 1 {
		return s
	}
	res := make(S, len(s))
	for i, e := range s {
		res[i] = f(e)
	}
	return res
}

// ChunkSlice chunks a slice in chunkSize big chunks and stores them in another slice
func ChunkSlice[S ~[]E, E ~[]X, X any](slice E, chunkSize int) (chunks S) {
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}
	return
}

// Insert elements into slice uniquely
func InsertUnique[S ~*[]E, E comparable](slice S, elem E) {
	for _, existing := range *slice {
		if existing == elem {
			return
		}
	}
	*slice = append(*slice, elem)
}

// Remove dups from slice.
func RemoveDups[S ~[]E, E comparable](slice S) (nodups S) {
	encountered := make(map[E]bool)
	for _, element := range slice {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

// Create slice of intersecting elements of 2 slices
func Intersection[S ~[]E, E comparable](s1, s2 S) (inter S) {
	hash := make(map[E]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = RemoveDups(inter)
	return
}
