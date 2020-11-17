package main 

import "fmt"


// func main (){
// 	number := []int{1,2,3,4,5,6,7,8,9,10}
// 	sum := 0;
// 	for i := 0; i<len(number); i++{
// 		sum += number[i]
// 	}
// 	fmt.Println("numbers",sum)
// }



// func main (){
// 	number := []int{1,2,3,4,5,6,7,8,9,10}
// 	sum := 0
// 	i := 0
// 	n := len(number)
// 	for i < n{
// 		sum += number[i]
// 		i++
// 	}
// 	fmt.Println(sum)
// }

// func main() {
// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// sum := 0
// i := 0
// n := len(numbers)
// for {
// sum += numbers[i]
// i++
// if i >= n {
// break
// }
// }
// fmt.Println("Sum is :: ", sum)
// }


// func main (){
// 	number := []int{1,2,3,4,5,6,7,8,9,10}

// 	sum := 0

// 	for i := 0; i<len(number); i++{
// 		 sum += number[i]
// 	}

// 	fmt.Println("sum is ",sum)
// }

/*func main() {
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sum := 0
i := 0
n := len(numbers)
for i < n {sum += numbers[i]
i++
}
fmt.Println("Sum is :: ", sum)
}*/

//  
// func main() {
// numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// sum := 0
// for _, val := range numbers {
// sum += val
// }
// fmt.Println("Sum is :: ", sum)
// }

func main() {
  tamil := []int{1,2,3,4,5,6,7,8,9,10}
   sum := 0
    for index,val := range tamil{
    	sum += val
    	fmt.Print("[", index, ",", val, "]")
    }
	fmt.Println("\nSum is :: ", sum)
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
	fmt.Println(k, " -> ", v)
 }
 str := "Hello, World!"
for index, c := range str {
fmt.Print("[", index, ",", string(c), "] ")
}

}