package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// run ./fetch https://www.geeksforgeeks.org/

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(urlCheck(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlCheck(url), err)
			os.Exit(1)
		}
		fmt.Printf("%d \nSTATUS CODE : %d",b,resp.StatusCode)
	}
		secs := time.Since(start).Seconds()
		time.Sleep(time.Second)
		output := fmt.Sprintf("execution fetching time : %.2f seconds\n", secs)
		output_file(output)
}

func urlCheck(url string) string{
	if !(strings.HasPrefix(url,"https://")) {
		real_url := "http://"+url
		url = real_url
	}else{
		return url
	}
	return url
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
