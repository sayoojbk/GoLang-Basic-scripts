package main

import(
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main(){

	b, err := ioutil.ReadFile("sample.txt")

	if err != nil{
		fmt.Println(err)
	}


	// convert bytes of file data into string data so as to be readble by humans.
	str: 
		string(b)

	fmt.Println(str)


	// To write to a file in GOLang

	file , err := os.Create("demo-sample.txt")

	if err !=nil{

		fmt.Println("We got an error" , err)

	}

	// If there is an error in the file creation close the file.
	defer file.close()

	fmt.Fprintf(file, "The file is created by GO Lang !")

}