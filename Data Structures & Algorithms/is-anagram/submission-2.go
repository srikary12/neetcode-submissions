import "maps"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sd := make(map[rune]int)
    td := make(map[rune]int)
    for _, char := range s {
        sd[char]++
    }
    for _, char := range t {
        td[char]++
    }
    return maps.Equal(sd, td)
}
