package main

import "fmt"

func main() {
	
	data := 10
	ptr := &data
	fmt.Println("Value stored at variable var is ", data)
	fmt.Println("Value stored at variable var is ", *ptr)
	fmt.Println("The address of variable var is ", &data)
	fmt.Println("The address of variable var is ", ptr)
}