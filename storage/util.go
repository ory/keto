package storage

func contains(target string, source []string) bool {
	for _, i := range source {
		if i == target {
			return true
		}
	}
	return false
}

func containsAll(targets []string, source []string) bool {
	for _, i := range targets {
		if !contains(i, source) {
			return false
		}
	}
	return true
}
