package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)


func main() {
	//ex: 3.10
	fmt.Println(comma([]string{"1","2","3"}))  // output: [1, 2, 3]
	
	//ex: 3.11
	fmt.Println(floating_point([]string{"3444","555","678","235"}))  // output: 34,44 55,5 67,8 23,5 

	//ex: 3.12
	fmt.Println(isAnagram("deniz","dinez"))  // output: false
	fmt.Println(isAnagram("selam","males"))  // output: true
}

func comma(str []string) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range str {
		if i > 0 {
			buf.WriteByte(',')
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%s", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func floating_point(number []string) string{
	var buf bytes.Buffer
	fmt.Println("Enter floating point : ")
	var point string
	fmt.Scan(&point)
	int_point,err := strconv.Atoi(point)
	if err !=nil{
		panic(err)
	}
	for i,_ := range number {
		a := len(number[i])
		b := strings.Split(number[i],"")
		if int_point < a {
			for j,v := range b {
				if j == int_point {				
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%s", v)
			}
		}else{
			fmt.Println("Enter floating point : ")
			var point string
			fmt.Scan(&point)
		}	
		buf.WriteByte(' ')
	}
	return buf.String()

} 

func isAnagram(str string, str2 string) bool {

	lenS := len(str)
	lenT := len(str2)
	b := strings.Split(str,"")
	c := strings.Split(str2,"")

	if lenS != lenT {
		return false
	}
	anagramMap := make(map[string]int)
	anagramMap2 := make(map[string]int)
	for i := 0; i < lenS; i++ {
		anagramMap[b[i]] = i
	}
	for i := lenT-1 ; i >= 0 ; i-- {
		anagramMap2[c[(lenT-1)-i]] = i
	}
	count := 0
	for i := 0; i < lenS; i++ {
		if anagramMap[b[i]] == anagramMap2[b[i]] {
			count++
			if count == lenS{
				return true
			}
		}
	}
	return false
}