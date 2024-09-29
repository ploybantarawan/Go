package main

import "fmt"

func main(){
	
	// x := 0
	// for x < 5 {
	// 	fmt.Println("the value of x is:", x)
	// 	x++
	// }

	// for i := 0; i<5; i++{
	// 	fmt.Println("the value of i is:", i)
	// }

	names := []string{"ploy", "meen", "ice", "fah"}

	// for i := 0; i<len(names); i++{
	// 	fmt.Println(names[i])
	// }

	for _, value := range names{
		 fmt.Printf("the values is %v \n", value)
	}

	fmt.Println(names)

} 