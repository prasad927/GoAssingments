package stack

import (
	"fmt"
)

type Stack struct {
   storage []interface{} //stack using array/slices
   top int
}

func (st *Stack) Push(val interface{}){
	
	if st.top==len(st.storage)-1 {
		oldStorage := st.storage
		newStorage := make([]interface{},2*len(st.storage)) //making new slice

		//copy data
		for i:=0;i<len(oldStorage);i++ {
			newStorage[i]=oldStorage[i]
		}

		st.storage = newStorage
	}

	st.top+=1
	st.storage[st.top]=val
}

func (st *Stack) Pop() (interface{},error){
    
	if st.top==-1 {
		return nil,fmt.Errorf("stack is empty")
	}

	valToReturn:=st.storage[st.top]
	st.top-=1

	return valToReturn,nil
}


func (st *Stack) Size() (int){
	return st.top+1
}

func (st *Stack) Peek() (interface{}){
	
	if st.top==-1 {
		return nil
	}

	return st.storage[st.top]
}

//stringer interface implemented..
func (st *Stack) String() (string){
	str:="["

	for i:=0;i<=st.top;i++ {
		str = str+" "+fmt.Sprint(st.storage[i])+""
	}
	return str+"] <--top"
}

//factory
func New(initialSize int) (*Stack){
	if initialSize<0 {
		panic("stack size cant be negative")
	}

	return &Stack{storage:make([]interface{},initialSize),top:-1}
}

