// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "math"

func isPerfectSquare(n int) bool {
	sq := math.Sqrt(float64(n))
	fmt.Println(n, sq)
	return sq == float64(int64(sq))
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func fermatFactor(n int) int {
	a := int(math.Ceil(math.Sqrt(float64(n))))
	fmt.Println(a)
	b2 := a*a - n

	for !isPerfectSquare(b2) && !isPrime(b2) {
		a += 1
		b2 = a*a - n
	}

	return a - int(math.Sqrt(float64(b2)))
}

func biggestPrimeFactor(n int) int {
	i := int(math.Ceil(math.Sqrt(float64(n))))
	for i > 0 {
		if n%i == 0 && isPrime(i) {
			return i
		}
		i--
	}

	return i
}

func main() {
	fmt.Println(biggestPrimeFactor(600851475143))
}
