package main

import "fmt"


func max(x, y int) int {
if x > y {
	fmt.Println("less")
return x
}
fmt.Println("greater")
return y
}
func main() {
fmt.Println(max(25, 20));
}
