// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	b := 0
	for i:= 0; i < 1000; i++ {
		if (i % 3 == 0 || i % 5 == 0) {
			b = b + i
		}
	}
	fmt.Println(b)
}
