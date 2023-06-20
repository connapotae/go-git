package main

import "fmt"

func sub(a, b int) int {
	return a - b

}

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b

}

func divide(a, b float64) float64 {
	return a / b
}

func main() {
	fmt.Println("Hello Go")
	fmt.Println(sum(10, 58))
	fmt.Println(sub(10, 20))
	fmt.Println(mul(10, 3))
	fmt.Println(divide(100, 20))
}
