package main

import "fmt"

func main(){

	var a =10
	var b =20 

	var c int

	c = add(a,b)
	fmt.Println("Addition of two numbers is : ", c)


}

// function prototype:   func function-name(param1, param2 variable type) return type


func add(num1, num2 int) int{


	var result int
	result = num1 + num2
	return result
}