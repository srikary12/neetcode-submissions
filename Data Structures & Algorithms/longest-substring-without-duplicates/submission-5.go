func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	res := 0
	charSet := make(map[byte]bool)
	for r = 0; r<len(s);r++ {
		for charSet[s[r]] {
			delete(charSet, s[l])
            l++
		}
		charSet[s[r]] = true
		if r - l + 1 > res {
            res = r - l + 1
        }
	}
	return res
}
