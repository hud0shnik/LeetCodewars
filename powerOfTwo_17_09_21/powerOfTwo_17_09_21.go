package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}

func main() {
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(-3))
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(999999999999999999))
	fmt.Println(isPowerOfTwo(1024))

}
