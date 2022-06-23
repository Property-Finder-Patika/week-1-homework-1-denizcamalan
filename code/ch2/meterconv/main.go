package main

import (
	"fmt"
	"os"
	"strconv"
)

type Kilometer float64
type Mile float64
//1kilometer (m) = 0.62137 miles (mi).

func main(){
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		k := Kilometer(t)
		m := Mile(t)
		
		output := fmt.Sprintf("%f = %s, %f = %s\n",k,KToM(k), m, MToK(m))
		output_file(output)
	}
}

func KToM(k Kilometer) Mile { return Mile(k*0.62137) }

func MToK(m Mile) Kilometer { return Kilometer(m/0.62137)}

func (k Kilometer) String() string { return fmt.Sprintf("%g km", k) }
func (m Mile) String() string { return fmt.Sprintf("%g mile", m) }

func create_file(){
	// create folder
	creating_doc,err := os.Create("output.txt")

	if err != nil {
		fmt.Println("done")
	}
	creating_doc.Close() 
}

func output_file(output string){

	writing_doc,err1 := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0666 )

	if err1 !=nil {
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