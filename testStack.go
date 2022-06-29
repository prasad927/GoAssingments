package main

import (
	stack "assingments/stack"
	"fmt"
)

func testStack(){
	fmt.Println("Test stack")

	myStack := stack.New(6)

	for i:=0;i<3;i++{
		fmt.Printf("The number %v is pushed \n",i)
		myStack.Push(i)
	}

	fmt.Printf("The current stack is %v \n",myStack)

	// newstack:=stack.New(2)
	// newstack.Pop()
}