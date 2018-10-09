package pagination

// Index uses limit, offset, and a slice's length to compute start and end indices for said slice.
func Index(limit, offset, length int) (start, end int) {
	if offset > length {
		return length, length
	} else if limit+offset > length {
		return offset, length
	}

	return offset, offset + limit
}
