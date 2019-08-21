package storage

func contains(target string, source []string) bool {
	for _, i := range source {
		if i == target {
			return true
		}
	}
	return false
}
