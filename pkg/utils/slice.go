package utils

// ChunkSlice chunks a slice in chunkSize big chunks and stores them in another slice
func ChunkSlice[S []E, E ~[]X, X any](slice E, chunkSize int) (chunks S) {
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
func InsertUnique[S ~[]E, E comparable](slice S, elem E) {
	for _, existing := range slice {
		if existing == elem {
			return
		}
	}
	slice = append(slice, elem)
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
