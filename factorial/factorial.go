package main

import "log"

func main() {
	log.Println("4! => ", factorial(4)) // 4*3*2*1
	log.Println("7! => ", factorial(7)) // 7*6*5*4*3*2*1
}

func factorial(n int) int {
	// 0! = 1
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
