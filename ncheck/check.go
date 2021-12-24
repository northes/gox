package ncheck

func StrInItem(val string, items []string) bool {
	for _, item := range items {
		if val == item {
			return true
		}
	}
	return false
}
