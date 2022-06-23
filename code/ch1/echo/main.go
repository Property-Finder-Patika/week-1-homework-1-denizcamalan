package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	output := fmt.Sprintln(strings.Join(os.Args[0:], " "))
	output_file(output)	
}

func create_file(){
	// create folder
	creating_doc,err := os.Create("output.txt")

	if err != nil {
		fmt.Println("done")
	}
	creating_doc.Close() 
}

func output_file(output string){

	writing_doc,err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0666 )
	errCheck(err)

	if err !=nil {
		create_file()
	} 
	
	k, err := writing_doc.WriteString(output) // k yazılan dosyanın boyutu
	errCheck(err)
	fmt.Println(k) 
}

func errCheck(err error){
	if err != nil{
		panic(err)
	}
}
