func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	for j := 0; j < len(t); j++ {
		hashMap[t[j]]--
	}
	for _, count := range hashMap {
		if count != 0 {
			return false
		}
	}
	return true

}
