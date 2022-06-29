package main

import (
	"fmt"
	"assingments/stack" //TYPO
)

func getPrecedance(operator interface{}) int{
     op:=operator.(string)

	 if op=="*" || op=="/"{
		 return 4
	 }else{
		 return 3
	 }

}

func evaluteExp(operands *stack.Stack,operator string){
	oprin2:=operands.Pop()
	oprin1:=operands.Pop()
    
	oprstr1,_:=oprin1.(string)
	oprstr2,_:=oprin2.(string)

	operands.Push(oprstr1+oprstr2+operator)
}

func InfixToPostfix(){
	var expression string
	fmt.Println("Enter the expression: ")
	fmt.Scan(&expression)
    
	//assuming inputs are valid infix expressions
    
	//stack vars are pointer types.
	operands_st:=stack.New(6)
	operators_st:=stack.New(6)

	for _,char:= range expression {
		strform := string(char)

		if strform=="+" || strform=="-" || strform=="*" || strform=="/" {

		     for operators_st.Size()!=0 && operators_st.Peek()!="(" && getPrecedance(operators_st.Peek())>=getPrecedance(strform){
				  operator:= operators_st.Pop()
				  strop,_:=operator.(string)
				  evaluteExp(operands_st,strop)
			 }

			 operators_st.Push(strform)      
		}else if strform=="("{
			 operators_st.Push(strform)
		}else if strform==")"{
			 for operators_st.Peek()!="("{
				operator:= operators_st.Pop()
				strop,_:=operator.(string)
				evaluteExp(operands_st,strop)
			 }
			 operators_st.Pop()
		}else if strform>="a" && strform<="z" || strform>="A" && strform<="Z"{
			// assuming a-z and A-Z in expression
			 operands_st.Push(strform)
		}

	}

	for operators_st.Size()!=0 {
        operator:= operators_st.Pop()
		strop,_:=operator.(string)
		evaluteExp(operands_st,strop)
	}

	fmt.Println("ANS: ",operands_st.Peek())
}