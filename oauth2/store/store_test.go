package store

import (
	"fmt"
	"testing"
)

var strLen = 16
var cases = 1000
var generated []string

func TestGenerateRandomString(t *testing.T) {
	generated = make([]string, cases)

	for i := 0; i < cases; i++ {
		t.Run(fmt.Sprintf("#%d", i), testGenerateRandomStringFunc(strLen, i))
	}
}

func testGenerateRandomStringFunc(n, i int) func(*testing.T) {
	return func(t *testing.T) {
		newStr := generateRandomString(n)

		for j := 0; j < i; j++ {
			if generated[j] == newStr {
				t.Errorf("Duplicate string on attempt #%d", i)
			}
		}

		generated[i] = newStr
	}
}