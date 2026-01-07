package hash

func canConstruct(ransomNote string, magazine string) bool {
	mag := [100001]int{}
	for _, val := range magazine {
		mag[val-'a']++
	}
	for _, val := range ransomNote {
		if mag[val-'a'] < 1 {
			return false
		}
		mag[val-'a']--
	}
	return true
}
