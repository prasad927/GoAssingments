package main

import (
	"fmt"
	"math"
)
func secondHighNum() (int,int){

	var size uint
	fmt.Println("Enter the size of array :")
	fmt.Scan(&size)
    
	//making slice to take size of arr from user/command line
	arr:=make([]int,size)
	fmt.Println("Enter integers nums in array ")

	var i uint
	for i=0;i<size;i++ {
		var val int
		fmt.Scan(&val)
		arr[i]=val
	}

	//find second max
	var max1 int=math.MinInt32
	var max2 int=math.MinInt32

    //assuming arr have unique elements only
	for _,num:=range arr {

		if num>max1 {
			max2=max1
			max1 = num
		}else if num>max2{
			max2 = num
		}
	}
	
	return max1,max2
}