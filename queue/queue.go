package queue

import "fmt"


type queueNode struct{
	data interface{}
	nxt  *queueNode //pointer of type queuenode
}

type queue struct {
	front *queueNode
	rear  *queueNode
	size uint
}

//enqueue
func (q *queue) Add(data interface{}){
	newnode:= &queueNode{
		data: data,
		nxt:nil,
	}

	if q.Size()==0 {
		q.front=newnode
		q.rear=newnode
	}else{
		q.rear.nxt = newnode
		q.rear = newnode
	}
	q.incsize()
}
//deque
func (q *queue) Remove(){
	if q.Size()==0{
		panic("queue is empty can't perform removal operation!")
	}

	if q.size==1{
		q.front=nil
		q.rear=nil
	}else{
		tempFront:=q.front
		q.front = q.front.nxt
		tempFront.nxt = nil
	}
	q.decsize()
}

func (q *queue) GetFront() interface{} {

	if q.Size()==0 {
		panic("queue is empty can't perform getFront operation!")
	}

	return q.front.data
}

func (q *queue) Size() uint{
	return q.size
}

func (q *queue) decsize(){
	q.size-=1
}

func (q *queue) incsize(){
	q.size+=1
}

//string representation of queue
func (q *queue) String() string {
	str:= "Front -> ["
	tempFront:=q.front

	for tempFront!=nil {
		str = str + fmt.Sprint(tempFront.data)+ " "
		tempFront = tempFront.nxt
	}

	str = str+ "] <-Rear"
	return str
}


func New() *queue{
	return &queue{}
}