package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	string_join()
	other_join()
}
var cases string

func string_join(){
	start := time.Now()
	cases := strings.Join(os.Args[1:],"")
	secs := time.Since(start).Seconds()
	time.Sleep(time.Second)
	fmt.Printf("execution time strings.Join : %.2fs  %s\n", secs, cases)
}
func other_join(){
	start := time.Now()
	for _, arg := range os.Args[1:] {
		cases = arg
	}
	secs := time.Since(start).Seconds()
	time.Sleep(time.Second)
	fmt.Printf("execution time other :%.2fs %s\n", secs, cases)
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
