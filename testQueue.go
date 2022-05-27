package main

import(
	Queue "assingments/queue"
	"fmt"
)


func testQueue(){
	q:=Queue.New()

	fmt.Println("Add to queue from 1 to 5")
	for i:=0;i<5;i++{
		q.Add(i+1)
	}

	fmt.Printf("The current queue: %v\n",q)

	fmt.Println("First Removal operation" )
	q.Remove()
	fmt.Printf("The current queue: %v\n",q)

	fmt.Printf("The current queue: %v\n",q)

	fmt.Println("Second Removal operation" )
	q.Remove()
	fmt.Printf("The current queue: %v\n",q)

	fmt.Printf("The current size: %v \n",q.Size())
	fmt.Printf("The current front: %v \n",q.GetFront())
}