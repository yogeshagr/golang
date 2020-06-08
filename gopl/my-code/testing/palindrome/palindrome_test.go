package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// BenchmarkIsPalindrome runs benchmark on IsPalindrome function
func benchmark(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func Benchmark10(b *testing.B) {
	benchmark(b, 10)
}

func Benchmark100(b *testing.B) {
	benchmark(b, 100)
}

func Benchmark1000(b *testing.B) {
	benchmark(b, 1000)
}
