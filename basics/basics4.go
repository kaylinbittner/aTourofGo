package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// int is the variable type
// if both arguments are the same type, only the last one requires a declaration
// i.e. func add(x, y int) int {



func main() {
	fmt.Println(add(42, 13))
}
