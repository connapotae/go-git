package main

import "fmt"

func sub(a, b int) int {
	return a - b

}

func sum(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Hello Go")
	fmt.Println(sum(10, 58))
	fmt.Println(sub(10, 20))
}
