package storage

func sliceContains(target, source []string) bool {
	m := make(map[string]int)

	for _, s := range source {
		m[s]++
	}

	for _, t := range target {
		if m[t] > 0 {
			return true
		}
	}

	return false
}

func contains(input []string, target string) bool {
	for _, i := range input {
		if i == target {
			return true
		}
	}
	return false
}
