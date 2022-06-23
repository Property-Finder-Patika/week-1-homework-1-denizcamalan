package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type result struct {
	counts    map[string]int
	fileNames map[string]string
}
var output []string

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)

	resultsEge := result{}

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Args[0], os.Stdin, resultsEge.counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts, fileNames)
			f.Close()
		}
	}
	i:=0
	for line, n := range counts {
		if n > 1 {
			output = append(output,fmt.Sprintf("%d\t%s\t%s\n",n, line, fileNames[line]))
			output_file(output[i])
			i++
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]int, fileNames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if strings.Contains(fileNames[input.Text()], arg) {
			continue
		} else {
			fileNames[input.Text()] = fileNames[input.Text()] + " " + arg
		}
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