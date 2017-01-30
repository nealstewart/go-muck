package main

import "fmt"

func divisibleUnder(n int, r int) bool {
	for i := 1; i <= r; i++ {
		if n%i != 0 {
			return false
		}
	}

	return true
}

func main() {
	i := 1
	for !divisibleUnder(i, 20) {
		i++
	}
	fmt.Println(i)
}
