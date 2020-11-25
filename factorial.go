package main

import "fmt"


func Factorial(i int) int {
// Termination Condition
if i <= 1 {
return 1
}// Body, Recursive Expansion
return i * Factorial(i-1)
}
func main(){
fmt.Println("factorial 5 is :: ", Factorial(5))
} 


