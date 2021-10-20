package main

import "fmt"

func main() {

	//x:= 100
	//fmt.Println(x)
	//someFunc()
	//someFunc2()
	//fmt.Println(miniGlobal)
	dialog()
}

func someFunc() {
	x := 150
	fmt.Println(x)
	fmt.Println(miniGlobal)
}

func someFunc2() {
	var test = "ok!"
	var point = 4.4

	fmt.Println(test)
	fmt.Println(point)
}

var miniGlobal = true

func dialog() {
	var x = "Hi"
	var y = "Hello"
	fmt.Println(x)
	{
		fmt.Println(y)
	}
}
