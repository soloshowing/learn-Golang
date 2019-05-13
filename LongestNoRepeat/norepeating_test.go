package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcabb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abcdef", 6},
		{"fjakhglkafgj", 7},
		{"房宝成房宝成", 3},
		{"", 0},
		{"b", 1},
		{"abcabcd", 4},
		{"黑化肥挥发发灰会花飞化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		if actual := longest(tt.s); actual != tt.ans {
			t.Errorf("longest(%s); got(%d); expected(%d)", tt.s, actual, tt.ans)
		}
	}
}

func BenchMarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞化肥挥发发黑会飞花"
	ans := 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := longest(s); actual != ans {
			b.Errorf("longest(%s); got(%d); expected(%d)", s, actual, ans)
		}
	}
}
