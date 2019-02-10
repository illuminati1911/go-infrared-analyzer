package utils

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"äabc", "cbaä"},
		{"Abc", "cbA"},
		{"AbC", "CbA"},
		{"TEST", "TSET"},
		{"010101010101", "101010101010"},
		{"saippuakauppias", "saippuakauppias"},
		{"...---&&&", "&&&---..."},
		{"This is a unit test.", ".tset tinu a si sihT"},
		{"Rock'n roll", "llor n'kcoR"},
		{"https://golang.org/", "/gro.gnalog//:sptth"},
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _, test := range tests {
		if r := Reverse(test.input); r != test.expected {
			t.Errorf("String reverse failure: Input: %s, Expected: %s, Received: %s", test.input, test.expected, r)
		}
	}
}

func TestSplitStringToChunks(t *testing.T) {
	var tests = []struct {
		s        string
		size     uint
		expected []string
	}{
		{"Hello", 2, []string{"He", "ll", "o"}},
		{"whats up 123", 5, []string{"whats", " up 1", "23"}},
		{" ", 500, []string{" "}},
		{"Hello", 1, []string{"H", "e", "l", "l", "o"}},
		{"Hello", 0, []string{"Hello"}},
		{"Hello", 12, []string{"Hello"}},
		{"abc123xyz", 3, []string{"abc", "123", "xyz"}},
		{"__**>>", 3, []string{"__*", "*>>"}},
		{"numbers98 7 ", 9, []string{"numbers98", " 7 "}},
		{"        ", 5, []string{"     ", "   "}},
		{"a", 0, []string{"a"}},
		{"", 55, []string{""}},
		{"", 0, []string{""}},
	}

	for _, test := range tests {
		splitted := SplitStringToChunks(test.s, test.size)
		if !isSlicesEqual(splitted, test.expected) {
			t.Errorf("String chunk split failure: Input string: %s, Input size: %d, Expected: %s, Received: %s",
				test.s, test.size, test.expected, splitted)
		}
	}
}

func isSlicesEqual(expected []string, splt []string) bool {
	if len(expected) != len(splt) {
		return false
	}
	for i, s := range expected {
		if s != splt[i] {
			return false
		}
	}
	return true
}
