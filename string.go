package main

import "fmt"

func main(){

	s := "tamil"
	r := []rune(s)
	r[0] = 'T'
	s2 :=string(r)
	
	fmt.Println("Before",s)
	
	fmt.Println("After",s2)
}