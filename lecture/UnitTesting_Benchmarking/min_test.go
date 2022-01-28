package utility

import (
	"fmt"
	"testing"
)

func TestGetMin(t *testing.T) {
	ans := GetMin(10, 3)
	if ans != 3 {
		t.Errorf("GetMin(10,3) = %v, want 3", ans)
	}
	ans = GetMin(3, 10)
	if ans != 3 {
		t.Errorf("GetMin(3,10) = %v, want 3", ans)
	}
}

func TestGetMinTable(t *testing.T) {
	var testCases = []struct {
		a, b, expected int
	}{
		{10, 3, 3},
		{100, 750, 100},
		{-5, -15, -15},
	}

	for _, test := range testCases {
		testName := fmt.Sprintf("%v,%v", test.a, test.b)
		testFunc := func(t *testing.T) {
			ans := GetMin(test.a, test.b)
			if ans != test.expected {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
		}
		t.Run(testName, testFunc)
	}

}

func BenchmarkGetMin(b *testing.B) {
	for n := 0; n < b.N; n++ { //b has a field N==billion
		GetMin(100, 53)
	}
}

//before running do:
//go mod init utility
//for testing:
//go test -v
//for bechmarking speed:
//go test -v -benchmark .
//for benchmarking memory:
//go test -v -benchmark . -benchmem
