package main
import "fmt"
func question9(){
	//find which number is not present in second array.
	arr1:= [...]int{1,2,3,4,5}
	arr2:= [...]int{2,3,1,0,5,0,0,0,6,5}

	track:=make(map[int]bool)
    

	for _,val:= range arr1 {
		track[val]=true
	}
    
    notPresentNumber:=make([]int,0)

	for _,val:=range arr2 {

		_,ispresent:=track[val]

		if !ispresent {
           notPresentNumber = append(notPresentNumber,val)
		   track[val]=true
		}
	}
    
	fmt.Println("next numbers are not present is second array : ",notPresentNumber)


}