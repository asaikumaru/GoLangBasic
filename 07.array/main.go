package main

import (
	"fmt"
)

func main() {

	var array [5]int

	array[0] = 100
	array[1] = 101
	array[2] = 102

	//fmt.Println(array)

	//for i:=0; i<len(array); i++ {
	//	fmt.Println(array[i])
	//}

	for i, v := range array {
		fmt.Printf("Index - %v, Value - %v\n", i, v)
	}

	var array2 = [...]string{"a", "b"}
	var array3 = [...]string{"c", "b"}
	fmt.Println(len(array2))

	fmt.Println(array2 == array3)

}
