package utils

// Reverse - reverses string. Runtime complexity O(n)
func Reverse(s string) string {
	rs := []rune(s)
	fs := ""
	for i := len(rs) - 1; i >= 0; i-- {
		fs += string(rs[i])
	}
	return fs
}

// SplitStringToChunks - splits string to chunks by size given.
func SplitStringToChunks(s string, size uint) []string {
	if uint(len(s)) <= size || size == 0 {
		return []string{s}
	}
	return append([]string{s[:size]}, SplitStringToChunks(s[size:], size)...)
}
