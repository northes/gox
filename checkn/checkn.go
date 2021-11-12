package checkn

func InArray(val string, items []string) bool {
	for _, item := range items {
		if val == item {
			return true
		}
	}
	return false
}
