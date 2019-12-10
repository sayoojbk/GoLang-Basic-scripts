package main

import (
	"fmt"
	"syscall"
)

func main(){

	err:= syscall.Mkdir("/demo" , 0754)

	if err == nil{
		fmt.Println("Directory created !")
	}


	cwd , _ := syscall.Getwd()

	fmt.Println("The current working directory is ", cwd)
	
}