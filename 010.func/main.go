package main

import "fmt"

func main() {

	//sum(5, 2)
	//fmt.Println(question("world"))
	//result, logic := someFunc(-1, 4, 323)
	//fmt.Println(result, logic)

	//result, _ := someFunc(-1, 4, 323)
	//fmt.Println(result)
	//someFuncTwo(10,55)
	fmt.Println(onlyTypes(3, "23")) // возможно использовать чтобы использовать для бага или таймаута сервера
}

func sum(a int, b int) {

	var z int
	z = a + b
	fmt.Println(z)
}
func question(q string) bool {
	if q == "question" {
		return true
	} else {
		return false
	}
}

func someFunc(a, b, c int) (int, bool) {
	if a > 0 && b > 5 {
		return 10, true
	} else if c > 100 {
		return 100, true
	} else {
		return 0, false
	}
}

func someFuncTwo(z, _ int) {
	fmt.Println(z)
}

func onlyTypes(int, string) bool {
	return true
}
