package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {
		fmt.Println(fizzbuzz(n))
	}
}

func fizzbuzz(n int) any {
	return []any{
		n, "Fizz", "Buzz", "FizzBuzz",
	}["300102100120100"[n%15]-'0']
}
