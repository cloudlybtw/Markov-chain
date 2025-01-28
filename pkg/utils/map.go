package utils

func MapContains(m map[string][]string, key string) bool {
	flag := false
	for mapkey, _ := range m {
		if mapkey == key {
			flag = true
		}
	}

	return flag
}
