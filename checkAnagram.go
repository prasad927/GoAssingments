package main

import "fmt"

func checkAnagram(){
  var str1 string
  var str2 string

  fmt.Println("Enter str1: ")
  fmt.Scan(&str1)

  fmt.Println("Enter str2: ")
  fmt.Scan(&str2)

    if len(str1)!=len(str2) {
	  fmt.Println("Given strings are not anagram")
	  return
	}

	var freqmap = make(map[string]uint)

	for _,char:=range str1 {
	   
		freq,ispresent:=freqmap[string(char)]

		if ispresent {
            freqmap[string(char)]=freq+1
		}else{
			freqmap[string(char)]=1
		}
	}
    
	var isAnagram bool = true
	for _,char:=range str2 {

		freq,isPresent:=freqmap[string(char)]

		if isPresent {
			if freq==1{
				 			//key,map
				delete(freqmap,string(char))
			}else{
				freqmap[string(char)]=freq-1
			}
		}else{
			isAnagram = false
			break
		}
	}


	if isAnagram {
		fmt.Println("The given string are anagrams")
	}else{
		fmt.Println("The given strings are not anagrams")
	}

}