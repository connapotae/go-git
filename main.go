package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b

}

func main() {
	fmt.Println("Hello Go")
	fmt.Println(sum(10, 58))
	fmt.Println(mul(10, 3))
}
