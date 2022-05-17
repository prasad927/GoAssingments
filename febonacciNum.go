package main

import "fmt"
func FebonacciNum(){
	var input int
	fmt.Println("How many febonacci nums you want to print:")
	fmt.Scan(&input)
    
	a:=0
	b:=1
	for i:=1;i<=input;i++ {
		fmt.Print(a," ")
        
		c:=a+b
		a=b
		b=c

	}
	fmt.Println()
}