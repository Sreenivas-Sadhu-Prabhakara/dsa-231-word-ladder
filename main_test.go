package main

import "testing"

func TestLadderLength(t *testing.T) {
	cases := []struct {
		begin, end string
		list       []string
		want       int
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0}, // end not in list
		{"a", "c", []string{"a", "b", "c"}, 2},
		{"hot", "dog", []string{"hot", "dog"}, 0}, // two letters apart, no bridge
		{"same", "same", []string{"same"}, 1},     // already there
		{"ab", "cd", []string{"ad", "cd"}, 3},     // ab -> ad -> cd
	}
	for _, c := range cases {
		if got := LadderLength(c.begin, c.end, c.list); got != c.want {
			t.Errorf("LadderLength(%q,%q) = %d, want %d", c.begin, c.end, got, c.want)
		}
	}
}
