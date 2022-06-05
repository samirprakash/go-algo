package main

import "log"

func main() {
	log.Println("fibonacci(2) => ", fibonacci(2)) // [0,1]
	log.Println("fibonacci(3) => ", fibonacci(3)) // [0,1,1]
	log.Println("fibonacci(5) => ", fibonacci(5)) // [0,1,1,2,3]
	log.Println("fibonacci(7) => ", fibonacci(7)) // [0,1,1,2,3,5,8]
}

func fibonacci(n int) []int {
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib
}
