package main

import "fmt"

func main(){

	var a [2]string

	a[0] = "dev"
	a[1] = "deb"

	fmt.Println(a[0])

	fmt.Println(a)

	var b [2]int

	b[0] =1
	b[1] =2
	fmt.Println(b)
}