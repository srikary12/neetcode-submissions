func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, l, maxf := 0, 0, 0

	for r := 0; r < len(s); r++ {
        count[s[r]]++
        if maxf<count[s[r]] {
            maxf = count[s[r]]
        }
        if (r-l+1) - maxf > k {
            count[s[l]]--
            l+=1
        }
        if res < (r-l+1) {
            res = r-l+1
        }
    }

    return res
}
