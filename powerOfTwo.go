package main

import "log"

func main() {
	log.Println("isPowerOfTwo(1) = ", isPowerOfTwo(1)) // true
	log.Println("isPowerOfTwo(2) = ", isPowerOfTwo(2)) // true
	log.Println("isPowerOfTwo(5) = ", isPowerOfTwo(5)) // false
}

func isPowerOfTwo(n int) bool {
	// bitwise operation
	return n > 0 && (n&(n-1)) == 0
}
