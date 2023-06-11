package main

import "fmt"

func sum_function(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Enter the first number: ")
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	sum := sum_function(a, b)
	fmt.Println(sum)

}
