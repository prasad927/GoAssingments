package main
import "fmt"


func longestPalindromInGivenString(){
     fmt.Println("Enter the string in which you want to find longest palidrome")
	 var inputstring string 
	 fmt.Scan(&inputstring)

	 if len(inputstring)==1 {
		fmt.Printf("The longest palindrome in given string is %v \n",inputstring) 
	 }

     longestPalindrome:=""
	 for i:=0;i<len(inputstring);i++ {
		 for j:=i+1;j<=len(inputstring);j++ {
			currsubstring:=inputstring[i:j]

			isPalindrome:=isStringIsPalindrome(currsubstring)

			if isPalindrome {
				if len(currsubstring) > len(longestPalindrome) {
					longestPalindrome = currsubstring
				}
			}
		 }
	 }

	 fmt.Printf("The longest palindrome in given string is : %v \n",longestPalindrome)
}