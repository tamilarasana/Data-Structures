 package main

import "fmt"


// func max(x, y int) int {
// if x > y {
// 	fmt.Println("less")
// return x
// }
// fmt.Println("greater")
// return y
// }
// func main() {
// fmt.Println(max(25, 20));
// }
// func IncrementPassByValue(x int) {
// 	x++
// 	}
// func main() {
// 	i := 10
// 	fmt.Println("Value of i before increment is : ", i)
// 	IncrementPassByValue(i)
// 	fmt.Println("Value of i after increment is : ", i)
// }


func IncrementPassByPointer(ptr *int) {
	(*ptr)++
	}
	func main() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is : ", i)
}