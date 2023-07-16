package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for n := 1; n <= 100; n++ {
		var want any

		switch {
		case n%3 == 0 && n%5 == 0:
			want = "FizzBuzz"
		case n%3 == 0:
			want = "Fizz"
		case n%5 == 0:
			want = "Buzz"
		default:
			want = n
		}

		got := fizzbuzz(n)

		if got != want {
			t.Errorf("fizzbuzz(%d) = %q, want %q", n, got, want)
		}
	}
}
