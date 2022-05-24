package main

import (
	"assingments/stack"
	"fmt"
)

func balancedParenthesis(){
	var expression string
	fmt.Println("Enter expression: ")
	fmt.Scan(&expression)



    stack_:=stack.New(5)
    

	for _,bracket:= range expression {
		strbrckt:=string(bracket)

		if strbrckt=="}" || strbrckt=="]" || strbrckt==")" {
			
			if stack_.Size()==0 {
				fmt.Println("Expression is not balanced")
				return
			}

			top:=stack_.Peek()
            
			if (strbrckt=="}" && top=="{") || (strbrckt=="]" && top=="[") || (strbrckt==")" && top=="(") {
				stack_.Pop()
			}else{
				fmt.Println("Expression is not balanced")
				return
			}

		}else{
			stack_.Push(strbrckt)
		}
	}

	if stack_.Size()==0 {
		fmt.Println("Expression is balanced")
	}else{
		fmt.Println("Expression is not balanced")
	}
}