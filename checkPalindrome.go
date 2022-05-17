package main

import "fmt"

func CheckNumOrStringIsPalindrome(){
     var choice int
	 fmt.Println("Which type you want to check for palindrome")
	 fmt.Println("1.string")
	 fmt.Println("2.number")
	 fmt.Println("Enter your choice")
	 fmt.Scan(&choice)
	
	 if choice==1 {
		 //string
            var input string 
		    fmt.Println("Enter string: ")
		    fmt.Scan(&input)
            

			isPalindrome:=isStringIsPalindrome(input)

			if isPalindrome {
				fmt.Println("The given string is palindrome")
			}else{
				fmt.Println("The given string is not palindrome")
			}

	 }else{
            var input int
		    fmt.Println("Enter number: ")
		    fmt.Scan(&input)

		    numForProcess:=input
		    reverseNumber:=0
		

		    for numForProcess!=0 {
              lastDigit:=numForProcess%10
			  numForProcess = numForProcess/10

			  reverseNumber = reverseNumber*10 + lastDigit
		    }


		    if(reverseNumber==input){
			  fmt.Println("The number is palindrome")
		    }else{
			  fmt.Println("The number is not palindrome")
		    }
	 } 
}


func isStringIsPalindrome(input string) bool{
	if(len(input)==1){
		return true
	}
	 
	low:=0
	high:=len(input)-1
	mid:=(len(input))/2

	for low<mid {
		if(input[low]!=input[high]){
			return false
	    }

		low++
		high--
	}

	return true
}