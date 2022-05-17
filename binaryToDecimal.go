package main
import "fmt"

func isValidInput(input string) bool{

	for _,data := range input {
		if string(data)=="0" || string(data)=="1"{ 
			continue
		}else{
			return false
		}
	}
	return true
}

func binaryToDecimal(){
    var binarynum string
	
	for {
		//process till user not enter valid number
	    fmt.Println("Enter binary number")
	    fmt.Scan(&binarynum)
	    isValid:=isValidInput(binarynum)

	    if !isValid {
		   fmt.Println("Cant move forward enter input number is not binary")
		   fmt.Println("Please enter again")
	    }else{
           break
	    }

    }

	ptr:=len(binarynum)-1
	valOfPowOfTwo:=1
	decimalNum:=0

	for ptr>=0 {
		data:=binarynum[ptr]
	    
		if string(data)=="1" {
			decimalNum = decimalNum + valOfPowOfTwo
		}

		valOfPowOfTwo = valOfPowOfTwo*2
		ptr--
	}

	fmt.Printf("The decimal equivalent of %v is %v \n",binarynum,decimalNum)


}