func isPalindrome(s string) bool {
    var builder strings.Builder
	for _, char := range s {
		// Check if the character is a letter or a number
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			// Convert to lowercase and write to the builder
			builder.WriteRune(unicode.ToLower(char))
		}
	}
    char := []rune(builder.String())
    l,r := 0, len(char) - 1
    for l <= r {
        if char[l] != char[r]{
            return false
        } else {
            l += 1
            r -= 1
        }
    }
    return true
}
