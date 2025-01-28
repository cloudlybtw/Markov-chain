package utils

func MapContains(m map[string][]string, key string) bool {
	flag := false
	for mapkey := range m {
		if mapkey == key {
			flag = true
		}
	}

	return flag
}
