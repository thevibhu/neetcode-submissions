func isPalindrome(s string) bool {
    cleaned := ""
    for _, ch := range strings.ToLower(s) {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            cleaned += string(ch)
        }
    }

  runes := []rune(cleaned)
    n := len(runes)
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}
