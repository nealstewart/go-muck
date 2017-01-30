// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func findBiggest(ceil int, floor int) int {
	lastFound := 0
	for i := ceil; i > floor; i-- {
		for j := ceil; j > floor; j-- {
			prod := i * j
			if isPalindrome(prod) && prod > lastFound {
				lastFound = prod
			}
		}
	}

	return lastFound
}

func main() {
	fmt.Println(findBiggest(999, 99))
}
