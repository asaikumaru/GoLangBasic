package main

import "fmt"

//const applicationName = "App Cosnt"
const (
	_ = iota // 0
	_ = iota
	_ = iota
	D = iota + 5
)

func main() {
	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	fmt.Println(D)
	//const number = 100
	//fmt.Println(applicationName)
	//fmt.Println(number)
	//
	//var number2 int = 10
	//result:= number + number2
	//fmt.Println(result)

	x, _, _ := TypesGo()
	fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)
}

func TypesGo() (int, bool, string) {
	return 100, false, "stringType"
}
