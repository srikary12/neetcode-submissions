func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[byte]int)
	l, res := 0,0
	for r:= 0;r < len(s); r++ {
		if idx, found := indexMap[s[r]]; found {
			l = max(idx+1,l)
		}
		indexMap[s[r]] = r
		res = max(res, r-l+1)
	}
	return res
}
