package main

import "fmt"

func main(){
	fmt.Println("Assignments...")
	var choice int
	fmt.Println("Enter your choice.")
	fmt.Println("1. CheckNumOrStringIsPalindrome")
	fmt.Println("2. FebonacciNum")
	fmt.Println("3. binaryToDecimal")
    fmt.Println("4. checkAnagram")
	fmt.Println("5. question9")
	fmt.Println("6. secondHighestNumber")
	fmt.Println("7. topTwoMaximumNumbers")
	fmt.Println("8. balancedParenthesis")
	fmt.Println("9. InfixToPostfix")
	fmt.Scan(&choice)
    

	switch choice {
     	 case 1:
			CheckNumOrStringIsPalindrome()
		 case 2:
			FebonacciNum()
		 case 3:
			 binaryToDecimal()
		 case 4:
			 checkAnagram()
		 case 5:
			 question9()
		 case 6:
			 _,secondHigh := secondHighNum()
			 fmt.Println("The second Highest number is: ",secondHigh)
		 case 7:
			 fmax,smax:=secondHighNum()
			 fmt.Println("The top two max numbers are: ",fmax,smax)
		 case 8:
			 balancedParenthesis()
		 case 9:
			 InfixToPostfix()
		 default:
			 fmt.Println("Wrong choice")
	}
}