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

func contains(target string, source []string) bool {
	for _, i := range source {
		if i == target {
			return true
		}
	}
	return false
}
