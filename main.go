package main

import "fmt"

func main(){

	age := 23
	name := "ploy"
	
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Println("hello ploy")
	fmt.Println("bye ploy")
	fmt.Println("my age is", age, "and my name is", name)

	//Printf (formatted strings) %_ = formatted specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 255.55)

	// Sprint (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}