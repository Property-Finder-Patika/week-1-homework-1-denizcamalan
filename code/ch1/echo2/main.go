// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	
	for _, arg := range os.Args[1:] {
		output := fmt.Sprintf("%s %s\n",os.Args[0],arg)
		output_file(output)
	}	
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