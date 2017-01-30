// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	secondLastTerm := 1
	lastTerm := 2

	sum := 2

	for lastTerm < 4000000 {
		tmp := lastTerm
		lastTerm = secondLastTerm + lastTerm
		secondLastTerm = tmp
		if lastTerm % 2 == 0 {
			sum += lastTerm
		}
		fmt.Println(lastTerm)
	}

	fmt.Println(sum)
}
