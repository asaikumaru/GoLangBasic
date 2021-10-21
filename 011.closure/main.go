package main

import "fmt"

func main() {

	//result := func(x,y int) (int) {		//// анонимная функция присвоенная переменной
	//	return x + y
	//}(10, 20)
	//
	//fmt.Println(result)
	//fmt.Println(sum(1,3))

	//closure("create a new text")

	closure()

	oneMore := func() int {
		z := closure() + 5
		return z
	}

	fmt.Println(oneMore())
}

func closure() int {

	t := 100

	result := func() {
		fmt.Println(t)
	}

	//z:= t + 10

	second := func(x int) {
		fmt.Println(t + x)
	}
	result()
	second(2)

	return t

}

//func closure(s string)  {
//
//	text := "sentence " + " - " + s
//
//	fmt.Println(text)
//
//	result := func() string {
//		fmt.Println(text + "some another text")		// замыкание
//		return text + "some another text"
//	}
//
//	s = "_abc"
//	sq := func() {
//		fmt.Println("ahaha" + s)			// замыкание
//	}
//	sq()
//	fmt.Println(result())
//}

//func sum(a, b int) int {
//
//	func(s string) {		// анонимная функция печати
//		fmt.Println(s)
//	}("some text")
//	return a + b
//}
